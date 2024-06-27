package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"

	// "gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/simple"

	"gonum.org/v1/gonum/graph/topo"
	"google.golang.org/grpc"

	// "gonum.org/v1/gonum/graph/multi"

	internalGraph "github.com/tcm5343/circular-dependency-detector/pkg/graph"
	pb "github.com/tcm5343/circular-dependency-detector/protos"
)

func main() {
	filePath := "/app/testing/alloy/dll.als"
	conn, err := grpc.Dial("172.17.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewFileStreamClient(conn)

	// Open the file to be sent
	// file, err := os.Open(filePath)
	// if err != nil {
	// 	log.Fatalf("failed to open file: %v", err)
	// }
	// defer file.Close()

	// Upload and analyze the file
	stream, err := client.UploadAndAnalyze(context.Background())
	if err != nil {
		log.Fatalf("could not upload and analyze file: %v", err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	buf := make([]byte, 1024)
	seq := int32(0)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read file: %v", err)
		}

		err = stream.Send(&pb.FileChunk{
			Content:  buf[:n],
			Sequence: seq,
		})
		if err != nil {
			log.Fatalf("could not send chunk: %v", err)
		}
		seq++
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("could not receive response: %v", err)
	}
	fmt.Printf("Analysis result: %s\n", response.GetResult())

	// Create a new multi-directed graph
	// g := multi.NewDirectedGraph()

	g := simple.NewDirectedGraph()

	// Parse the DOT data into the graph
	err = dot.Unmarshal([]byte(response.GetResult()), g)
	if err != nil {
		log.Fatalf("failed to unmarshal DOT data: %v", err)
	}

	// Print out the nodes and edges
	fmt.Println("Nodes:")
	for nodes := g.Nodes(); nodes.Next(); {
		node := nodes.Node()
		fmt.Printf("Node ID: %v\n", node.ID())
	}

	fmt.Println("Edges:")
	for edges := g.Edges(); edges.Next(); {
		edge := edges.Edge()
		fmt.Printf("Edge from %v to %v\n", edge.From().ID(), edge.To().ID())
	}

	// ... test code ...

	// todo: maybe print out the args
	inputGraphPath := flag.String("input-file", "SOME DEFAULT PATH", "absolute path to graph input file")
	failOnCycle := flag.Bool("fail-on-cycle", true, "flag if action should fail if a cycle is detected")
	flag.Parse()

	slog.SetLogLoggerLevel(slog.LevelDebug)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// var inputGraphPath string
	// inputFile, present := os.LookupEnv("INPUT_FILE")
	// if present {
	// 	inputGraphPath = inputFile
	// } else {
	// 	inputGraphPath = argsWithoutProg[0]
	// }

	fp, err := os.Open(*inputGraphPath) // fixme: this path can't contain spaces for some reason...
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	lg, err := internalGraph.ParseInputGraph(fp)
	if err != nil {
		panic(err)
	}

	// cycles := topo.DirectedCyclesIn(lg.Graph)
	cycles := topo.DirectedCyclesIn(g)
	slog.Debug("directed cycles in", "count", len(cycles), "cycles", cycles)

	if len(cycles) > 0 && *failOnCycle {
		slog.Error("failing due to cycles being found")
		os.Exit(1)
	}

	if len(cycles) > 0 {
		slog.Info("skipping topological generations . . .")
	} else {
		tg, err := internalGraph.TopologicalGenerationsOf(lg.Graph)
		if err != nil {
			panic(err)
		}
		slog.Info("topological generations", "generations", tg)
	}
}

package main

import (
	"flag"
	"log"
	"log/slog"
	"os"

	"gonum.org/v1/gonum/graph/topo"

	internalGraph "github.com/tcm5343/circular-dependency-detector/pkg/graph"
)

func main() {

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

	cycles := topo.DirectedCyclesIn(lg.Graph)
	// cycles := topo.DirectedCyclesIn(g)
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

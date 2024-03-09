package main

import (
	"log"
	"log/slog"
	"os"

	"gonum.org/v1/gonum/graph/topo"

	"github.com/tcm5343/circular-dependency-detector/pkg/graph"
)

func main() {
	argsWithoutProg := os.Args[1:]

	slog.SetLogLoggerLevel(slog.LevelInfo)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// fmt.Println(argsWithoutProg)
	// fmt.Println("/app/" + inputGraphPath)
	var inputGraphPath string
	val, present := os.LookupEnv("INPUT_FILE")
	if present {
		inputGraphPath = val
	} else {
		inputGraphPath = argsWithoutProg[0]
	}

	fp, err := os.Open(inputGraphPath) // fix: this path can't contain spaces for some reason...
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	lg, err := graph.ParseInputGraph(fp)
	if err != nil {
		panic(err)
	}

	cycles := topo.DirectedCyclesIn(lg.Graph)
	slog.Debug("directed cycles in", "count", len(cycles), "cycles", cycles)

	if len(cycles) > 0 {
		slog.Info("skipping topological generations . . .")
	} else {
		tg, err := graph.TopologicalGenerationsOf(lg.Graph)
		if err != nil {
			panic(err)
		}
		slog.Info("topological generations", "generations", tg)
	}
}

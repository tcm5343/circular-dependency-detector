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
	inputGraphPath := argsWithoutProg[0]

	slog.SetLogLoggerLevel(slog.LevelInfo)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// fmt.Println(argsWithoutProg)
	// fmt.Println("/app/" + inputGraphPath)

	fp, err := os.Open("/app/" + inputGraphPath) // fix: this path can't contain spaces for some reason...
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	wg, err := graph.BuildDirectedGraph(fp)
	if err != nil {
		panic(err)
	}

	cycles := topo.DirectedCyclesIn(wg.Graph)
	slog.Debug("directed cycles in", "count", len(cycles), "cycles", cycles)

	if len(cycles) > 0 {
		slog.Info("skipping topological generations . . .")
	} else {
		tg, err := graph.TopologicalGenerationsOf(wg.Graph)
		if err != nil {
			panic(err)
		}
		slog.Info("topological generations", "generations", tg)
	}
}

package main

import (
	"fmt"

	"gonum.org/v1/gonum/graph/topo"

	"github.com/tcm5343/circular-dependency-detector/pkg/graph"
)

func main() {
	dg, err := graph.BuildGraph("./list/adjacency_list_cycle_2.csv")
	if err != nil {
		panic(err)
	}

	cycles := topo.DirectedCyclesIn(dg)
	fmt.Println(cycles)
}

package main

import (
	"fmt"
	"strconv"

	"gonum.org/v1/gonum/graph/topo"

	"github.com/tcm5343/circular-dependency-detector/pkg/graph"
)

func main() {
	dg, err := graph.BuildDirectedGraph("/app/testing/adj_list_cycle.txt")
	if err != nil {
		panic(err)
	}

	cycles := topo.DirectedCyclesIn(dg)
	numberOfCycles := strconv.Itoa(len(cycles))

	fmt.Println(numberOfCycles + " cycle(s) identified")

	if len(cycles) > 0 {
		fmt.Println(cycles)
		fmt.Println("Skipping topological generation . . .")
	} else {
		tg, err := graph.TopologicalGenerationsOf(dg)
		if err != nil {
			panic(err)
		}

		fmt.Print("topological generations: ")
		fmt.Println(tg)
	}
}

package main

import (
	"fmt"
	"os"
	"strconv"

	"gonum.org/v1/gonum/graph/topo"

	"github.com/tcm5343/circular-dependency-detector/pkg/graph"
)

func main() {
	argsWithoutProg := os.Args[1:]
	inputGraphPath := argsWithoutProg[0]

	// fmt.Println(argsWithoutProg)
	// fmt.Println("/app/" + inputGraphPath)

	dg, err := graph.BuildDirectedGraph("/app/" + inputGraphPath) // fix: this path can't contain spaces for some reason...
	if err != nil {
		panic(err)
	}

	cycles := topo.DirectedCyclesIn(dg)
	numberOfCycles := strconv.Itoa(len(cycles))

	fmt.Println(numberOfCycles + " cycle(s) identified")

	if len(cycles) > 0 {
		fmt.Println(cycles)
		fmt.Println("skipping topological generation . . .")
	} else {
		tg, err := graph.TopologicalGenerationsOf(dg)
		if err != nil {
			panic(err)
		}

		fmt.Print("topological generations: ")
		fmt.Println(tg)
	}
}

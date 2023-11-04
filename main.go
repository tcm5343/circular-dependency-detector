package main

import (
	"github.com/tcm5343/circular-dependency-detector/pkg/graph"
)

func main() {
	graph.BuildGraph("./list/adjacency_list_cycle_2.csv")
}

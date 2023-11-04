package graph

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"gonum.org/v1/gonum/graph/simple"
)

func BuildDirectedGraph(filePath string) (*simple.DirectedGraph, error) {
	fp, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	reader := csv.NewReader(fp)
	reader.FieldsPerRecord = -1
	reader.Comma = ' '
	dg := simple.NewDirectedGraph()

	lines, err := reader.ReadAll()
	if err != nil {
		return dg, fmt.Errorf("error reading adjacency list: %w", err)
	}

	for _, line := range lines {
		idFrom, _ := strconv.ParseInt(line[0], 10, 64)
		for i := 1; i < len(line); i++ {
			idTo, _ := strconv.ParseInt(line[i], 10, 64)
			dg.SetEdge(simple.Edge{F: simple.Node(idFrom), T: simple.Node(idTo)})
		}
	}

	return dg, nil
}

func TopologicalGenerationsOf(dg *simple.DirectedGraph) (*simple.DirectedGraph, error) {
	return dg, nil
}

package graph

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"gonum.org/v1/gonum/graph/simple"
)

func BuildGraph(filePath string) error {
	fp, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	reader := csv.NewReader(fp)
	reader.FieldsPerRecord = -1
	reader.Comma = ' '

	lines, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("error reading adjacency list: %w", err)
	}

	dg := simple.NewDirectedGraph()
	for _, line := range lines {
		idFrom, _ := strconv.ParseInt(line[0], 10, 64)
		for i := 1; i < len(line); i++ {
			idTo, _ := strconv.ParseInt(line[i], 10, 64)
			dg.SetEdge(simple.Edge{F: simple.Node(idFrom), T: simple.Node(idTo)})
		}
	}

	return nil
}

package graph

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

func BuildGraph(list string) error {
	// Define Graph
	dg := simple.NewDirectedGraph()
	// Open File
	fp, err := os.Open(list)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	// Read rows and add edges
	reader := csv.NewReader(fp)
	reader.FieldsPerRecord = -1
	reader.Comma = ' '
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("error reading adjacency list: %w", err)
		}

		idFrom, _ := strconv.ParseInt(record[0], 10, 64)
		checkForNode := dg.Node(idFrom)

		if checkForNode == nil {
			dg.AddNode(simple.Node(idFrom))
		}
		for i := 1; i < len(record); i++ {
			idTo, _ := strconv.ParseInt(record[i], 10, 64)
			checkForNode = dg.Node(idTo)
			if checkForNode == nil {
				dg.AddNode(simple.Node(idTo))
			}
			edge := dg.NewEdge(dg.Node(idFrom), dg.Node(idTo))
			dg.SetEdge(edge)
		}

	}

	j := topo.DirectedCyclesIn(dg)
	fmt.Println(j)

	return nil
}

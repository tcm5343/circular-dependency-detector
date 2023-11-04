package graph

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	// "gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/multi"
)

func BuildDirectedGraph(filePath string) (*multi.DirectedGraph, error) {
	fp, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	reader := csv.NewReader(fp)
	reader.FieldsPerRecord = -1
	reader.Comma = ' '
	dg := multi.NewDirectedGraph()

	lines, err := reader.ReadAll()
	if err != nil {
		return dg, fmt.Errorf("error reading adjacency list: %w", err)
	}

	for _, line := range lines {
		idFrom, _ := strconv.ParseInt(line[0], 10, 64)
		idTo, _ := strconv.ParseInt(line[1], 10, 64)
		dg.SetLine(dg.NewLine(multi.Node(idFrom), multi.Node(idTo)))
	}

	// fmt.Println(strconv.Itoa(dg.Nodes().Len()) + " node(s) and " + strconv.Itoa(len(lines)) + " edge(s) identified")

	return dg, nil
}

func TopologicalGenerationsOf(dg *multi.DirectedGraph) ([][]int, error) {
	indegreeMap := make(map[int]int)
	var zeroIndegree []int
	var generations [][]int

	nodes := dg.Nodes()
	for nodes.Next() {
		node := nodes.Node()
		inNodes := dg.To(node.ID())
		if inNodes.Len() > 0 {
			indegreeMap[int(node.ID())] = inNodes.Len()
		} else {
			zeroIndegree = append(zeroIndegree, int(node.ID()))
		}
		// fmt.Println(strconv.Itoa(int(node.ID())) + " has an indegree of " + strconv.Itoa(inNodes.Len()))
	}

	for zeroIndegree != nil {
		generations = append(generations, zeroIndegree)
		zeroIndegree = nil

		for _, nodeId := range generations[len(generations)-1] {
			outNodes := dg.From(int64(nodeId))
			for outNodes.Next() {
				node := outNodes.Node()
				outNodeId := int(node.ID())
				indegreeMap[outNodeId] -= 1
				if indegreeMap[int(node.ID())] == 0 {
					// fmt.Println("Adding " + strconv.Itoa(outNodeId) + " to zeroIndegree slice")
					zeroIndegree = append(zeroIndegree, outNodeId)
					delete(indegreeMap, outNodeId)
				}
			}
		}
	}

	return generations, nil
}

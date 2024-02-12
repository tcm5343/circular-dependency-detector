package graph

import (
	"encoding/csv"
	"fmt"
	"io"
	"log/slog"
	"strings"

	"gonum.org/v1/gonum/graph/multi"
)

// func (g LabledGraph) nodeFor(word string) graph.Node {
// 	id, ok := g.labels[word]
// 	if !ok {
// 		return nil
// 	}
// 	return g.UndirectedGraph.Node(id)
// }

// // node is a word node in a LabledGraph.
// type node struct {
// 	word string
// 	id   int64
// }

// func (n node) ID() int64      { return n.id }
// func (n node) String() string { return n.word }

type LabledGraph struct {
	labels map[string]int64
	Graph  *multi.DirectedGraph
}

func newLabledGraph() LabledGraph {
	return LabledGraph{
		labels: make(map[string]int64),
		Graph:  multi.NewDirectedGraph(),
	}
}

func (g *LabledGraph) include(idFrom string, idsTo []string) {
	// handle source node
	if _, exists := g.labels[idFrom]; !exists {
		u := g.Graph.NewNode()
		g.Graph.AddNode(u)
		g.labels[idFrom] = u.ID()
	}
	fromNode, _ := g.Graph.NodeWithID(g.labels[idFrom])

	// handle destination nodes
	for _, element := range idsTo {
		if _, exists := g.labels[element]; !exists {
			toNode := g.Graph.NewNode()
			g.Graph.AddNode(toNode)
			g.labels[element] = toNode.ID()
			g.Graph.SetLine(g.Graph.NewLine(fromNode, toNode))
		}
		slog.Debug("adding edge", "sourceNode", idFrom, "destinationNode", element)
		g.Graph.SetLine(g.Graph.NewLine(fromNode, g.Graph.Node(g.labels[element])))
	}
}

func removeElementsAfterPrefix(slice []string, prefix string) []string {
	// If a string in the slice starts with the prefix, remove it and all elements to the right of it
	for i, str := range slice {
		if strings.HasPrefix(str, prefix) {
			return slice[:i]
		}
	}
	return slice
}

func BuildDirectedGraph(inputFile io.Reader) (*LabledGraph, error) {
	delimiter := ' '     // todo: make this a parameter
	commentMarker := "#" // todo: make this a parameter

	reader := csv.NewReader(inputFile)
	reader.FieldsPerRecord = -1
	reader.Comma = delimiter

	wg := newLabledGraph()

	lines, err := reader.ReadAll()
	if err != nil {
		return &wg, fmt.Errorf("error reading adjacency list: %w", err)
	}

	for _, line := range lines {
		if !strings.HasPrefix(line[0], commentMarker) { // is entire line a comment
			wg.include(line[0], removeElementsAfterPrefix(line[1:], commentMarker)) // todo: this is inefficient since we loop over toNodes twice, improve
		}
	}
	return &wg, nil
}

func TopologicalGenerationsOf(dg *multi.DirectedGraph) ([][]int, error) {
	inDegreeMap := make(map[int]int)
	var zeroInDegree []int
	var generations [][]int

	nodes := dg.Nodes()
	for nodes.Next() {
		node := nodes.Node()
		inNodes := dg.To(node.ID())
		if inNodes.Len() > 0 {
			inDegreeMap[int(node.ID())] = inNodes.Len()
		} else {
			zeroInDegree = append(zeroInDegree, int(node.ID()))
		}
		// fmt.Println(strconv.Itoa(int(node.ID())) + " has an in-degree of " + strconv.Itoa(inNodes.Len()))
	}

	for zeroInDegree != nil {
		generations = append(generations, zeroInDegree)
		zeroInDegree = nil

		for _, nodeId := range generations[len(generations)-1] {
			outNodes := dg.From(int64(nodeId))
			for outNodes.Next() {
				node := outNodes.Node()
				outNodeId := int(node.ID())
				inDegreeMap[outNodeId] -= 1
				if inDegreeMap[int(node.ID())] == 0 {
					// fmt.Println("Adding " + strconv.Itoa(outNodeId) + " to zeroIndegree slice")
					zeroInDegree = append(zeroInDegree, outNodeId)
					delete(inDegreeMap, outNodeId)
				}
			}
		}
	}

	return generations, nil
}

package graph

import (
	"encoding/csv"
	"fmt"
	"io"
	"log/slog"
	"strings"

	"gonum.org/v1/gonum/graph/multi"
)

// func (g WordGraph) nodeFor(word string) graph.Node {
// 	id, ok := g.ids[word]
// 	if !ok {
// 		return nil
// 	}
// 	return g.UndirectedGraph.Node(id)
// }

// // node is a word node in a WordGraph.
// type node struct {
// 	word string
// 	id   int64
// }

// func (n node) ID() int64      { return n.id }
// func (n node) String() string { return n.word }

type WordGraph struct {
	ids   map[string]int64
	Graph *multi.DirectedGraph
}

func newWordGraph() WordGraph {
	return WordGraph{
		ids:   make(map[string]int64),
		Graph: multi.NewDirectedGraph(),
	}
}

func (g *WordGraph) include(idFrom string, idsTo []string) {
	// handle source node
	if _, exists := g.ids[idFrom]; !exists {
		u := g.Graph.NewNode()
		g.Graph.AddNode(u)
		g.ids[idFrom] = u.ID()
	}
	fromNode, _ := g.Graph.NodeWithID(g.ids[idFrom])

	// handle destination nodes
	for _, element := range idsTo {
		if _, exists := g.ids[element]; !exists {
			toNode := g.Graph.NewNode()
			g.Graph.AddNode(toNode)
			g.ids[element] = toNode.ID()
			g.Graph.SetLine(g.Graph.NewLine(fromNode, toNode))
		}
		slog.Debug("adding edge", "sourceNode", idFrom, "destinationNode", element)
		g.Graph.SetLine(g.Graph.NewLine(fromNode, g.Graph.Node(g.ids[element])))
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

func BuildDirectedGraph(inputFile io.Reader) (*WordGraph, error) {
	delimiter := ' '     // todo: make this a parameter
	commentMarker := "#" // todo: make this a parameter

	reader := csv.NewReader(inputFile)
	reader.FieldsPerRecord = -1
	reader.Comma = delimiter

	wg := newWordGraph()

	lines, err := reader.ReadAll()
	if err != nil {
		return &wg, fmt.Errorf("error reading adjacency list: %w", err)
	}

	for _, line := range lines {
		if !strings.HasPrefix(line[0], commentMarker) { // check if entire line is a comment
			wg.include(line[0], removeElementsAfterPrefix(line[1:], commentMarker)) // todo: this is inefficient since we loop over toNodes twice, improve
		} else {
			// fmt.Printf("skipping %v due to entire line being a comment\n", line)
		}
	}
	return &wg, nil
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

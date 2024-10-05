package graph

import (
	"encoding/csv"
	"fmt"
	"io"
	"log/slog"
	"strings"

	"gonum.org/v1/gonum/graph/multi"
)

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

func ParseInputGraph(inputFile io.Reader) (*LabledGraph, error) {
	// todo: set the delimiter and commentMarker via .env config and pass into function
	delimiter := ' '     // todo: make this a parameter
	commentMarker := "#" // todo: make this a parameter

	// todo: the delimiter may not be the same as the comment marker
	// todo: no node name may contain the delimiter or commentMarker

	reader := csv.NewReader(inputFile)
	reader.FieldsPerRecord = -1
	reader.Comma = delimiter

	lg := newLabledGraph()

	lines, err := reader.ReadAll()
	if err != nil {
		return &lg, fmt.Errorf("error reading adjacency list input: %w", err)
	}

	for _, line := range lines {
		// todo: check for empty lines too
		if !strings.HasPrefix(line[0], commentMarker) { // is entire line a comment
			lg.include(line[0], removeElementsAfterPrefix(line[1:], commentMarker)) // todo: this is inefficient since we loop over toNodes twice, improve
		}
	}
	return &lg, nil
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

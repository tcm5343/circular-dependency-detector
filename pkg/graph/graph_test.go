package graph

import (
    "testing"
    "github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/graph/multi"
)

func TestBuildDirectedGraphNoCycle(t *testing.T) {
    expGraph := multi.NewDirectedGraph()
	expGraph.SetLine(expGraph.NewLine(multi.Node(0), multi.Node(1)))
	expGraph.SetLine(expGraph.NewLine(multi.Node(0), multi.Node(2)))
	expGraph.SetLine(expGraph.NewLine(multi.Node(1), multi.Node(3)))
	expGraph.SetLine(expGraph.NewLine(multi.Node(1), multi.Node(4)))
	expGraph.SetLine(expGraph.NewLine(multi.Node(2), multi.Node(5)))
	expGraph.SetLine(expGraph.NewLine(multi.Node(2), multi.Node(6)))
	expGraph.SetLine(expGraph.NewLine(multi.Node(3), multi.Node(4)))
	actGraph, _ := BuildDirectedGraph("../../list/adjacency_list.csv")
	assert.Equal(t, expGraph, actGraph)
}
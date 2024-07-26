package dot

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
	"gonum.org/v1/gonum/graph/multi"
)

type DotMultiGraph struct {
	*multi.DirectedGraph
}

func NewDotMultiGraph() *DotMultiGraph {
	return &DotMultiGraph{DirectedGraph: multi.NewDirectedGraph()}
}

func (g *DotMultiGraph) AddEdge(from, to graph.Node) graph.Line {
	e := &DotLine{Line: g.DirectedGraph.NewLine(from, to), attrs: make(map[string]string)}
	return e
}

func (g *DotMultiGraph) NewLine(from, to graph.Node) graph.Line {
	line := g.DirectedGraph.NewLine(from, to)
	dotLine := &DotLine{Line: line, attrs: make(map[string]string)}
	return dotLine
}

func (g *DotMultiGraph) SetLine(e graph.Line) {
	e.(*DotLine).g = g
	g.DirectedGraph.SetLine(e)
}

type DotLine struct {
	g *DotMultiGraph
	graph.Line
	attrs map[string]string
}

// func (e *dotLine) ReversedLine() *dotLine {
// 	reversedLine := &dotLine{
// 		g:     e.g,
// 		Line:  e.g.DirectedGraph.NewLine(e.To(), e.From()),
// 		attrs: make(map[string]string),
// 	}
// 	for k, v := range e.attrs {
// 		reversedLine.attrs[k] = v
// 	}
// 	return reversedLine
// }

func (e *DotLine) SetAttribute(attr encoding.Attribute) error {
	// fmt.Printf("attr: %s\n", attr)

	if attr.Key == "dir" && attr.Value == "both" {
		// todo: method to reverse a dotLine instead
		reversedLine := e.ReversedLine()
		reversedDotLine := &DotLine{
			g:     e.g,
			Line:  reversedLine,
			attrs: make(map[string]string),
		}
		for k, v := range e.attrs {
			reversedDotLine.attrs[k] = v
		}
		e.g.DirectedGraph.SetLine(reversedDotLine)
	}

	e.attrs[attr.Key] = attr.Value
	return nil
}

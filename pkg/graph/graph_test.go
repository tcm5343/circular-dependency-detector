package graph

import (
	"io"
	"log"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/topo"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	dotGraph "github.com/tcm5343/circular-dependency-detector/pkg/dot"
	"github.com/tcm5343/circular-dependency-detector/pkg/rpc"
)

type testEdge = struct {
	from     string
	to       string
	expected bool
}

// func TestTopologicalGenerationsOfCyclicGraph(t *testing.T) {
// 	fmt.Println("TestTopologicalGenerationsOfCyclicGraph")
// }

// func TestTopologicalGenerationsOfAcyclicGraph(t *testing.T) {
// 	fmt.Println("TestTopologicalGenerationsOfAcyclicGraph")
// }

// func TestTopologicalGenerationsOfMultiGraph(t *testing.T) {
// 	fmt.Println("TestTopologicalGenerationsOfMultiGraph")
// }

// func TestTopologicalGenerationsOfEmptyGraph(t *testing.T) {
// 	fmt.Println("TestTopologicalGenerationsOfEmptyGraph")
// }

// func TestBuildDirectedGraphFileNotFoundError(t *testing.T) {
// 	fmt.Println("TestBuildDirectedGraphFileNotFoundError")
// }

func trimWhitespaceFromLines(input string) string { // todo: move to internal/testing.go and unit test
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimLeft(line, " \t") // trim leading spaces and tabs
	}
	return strings.Join(lines, "\n")
}

func areValuesDistinct(m map[string]int64) bool { // todo: move to internal/testing.go and unit test
	seen := make(map[int64]struct{})
	for _, v := range m {
		if _, exists := seen[v]; exists {
			return false // If the value has been seen before, not distinct
		}
		seen[v] = struct{}{}
	}
	return true // All values are distinct
}

// func TestParsingGraphs(t *testing.T) {
// 	ug := `
// 	digraph "graph" {
// 		graph [fontsize=12]
// 		node [fontsize=12]
// 		edge [fontsize=12]
// 		rankdir=TB;
// 		"N0" -> "N1" [uuid = "<Node1, Node0>", color = "#000000", fontcolor = "#000000", style = "solid", label = "edges", dir = "forward"]
// 		"N0" -> "N2" [uuid = "<Node1, Node2>", color = "#000000", fontcolor = "#000000", style = "solid", label = "edges", dir = "both"]
// 		"N2" -> "N1" [uuid = "<Node2, Node0>", color = "#000000", fontcolor = "#000000", style = "solid", label = "edges", dir = "forward"]
// 		"N0" [uuid="Node1", label="Node1", color="#ffd700", fontcolor = "#000000", shape = "box", style = "filled, solid"]
// 		"N1" [uuid="Node0", label="Node0", color="#ffd700", fontcolor = "#000000", shape = "box", style = "filled, solid"]
// 		"N2" [uuid="Node2", label="Node2\n($cyclic_n)", color="#ffd700", fontcolor = "#000000", shape = "box", style = "filled, solid"]
// 	}
// 	`

// 	g := newDotMultiGraph()
// 	err := dot.UnmarshalMulti([]byte(ug), g)
// 	if err != nil {
// 		log.Fatalf("failed to unmarshal DOT data: %v", err)
// 	}

// 	// fmt.Println("Edges:")
// 	// for edges := g.Edges(); edges.Next(); {
// 	// 	edge := edges.Edge()
// 	// 	fmt.Printf("Edge from %v to %v\n", edge.From().ID(), edge.To().ID())
// 	// 	for lines := g.Lines(edge.From().ID(), edge.To().ID()); lines.Next(); {
// 	// 		line := lines.Line().(*DotLine)
// 	// 		fmt.Printf("Attributes: %v\n", line.attrs)
// 	// 	}
// 	// }

// 	data, err := dot.Marshal(g, "Example Graph", "", "\t")
// 	if err != nil {
// 		log.Fatalf("error marshalling graph to DOT: %v", err)
// 	}
// 	t.Log(string(data))
// }

func TestAlloyCyclicGraphs(t *testing.T) {
	filePath := "/app/testing/alloy/directed_graph.als"
	alloyCommand := "run cyclic for 3"

	client, conn, err := rpc.NewClient("172.17.0.1:8080")
	if err != nil {
		t.Fatalf("could not create connection client to server: %v", err)
	}
	defer conn.Close()

	stream, err := rpc.UploadFile(client, filePath, alloyCommand)
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			t.Logf("gRPC error code: %s\n", s.Code())
			t.Logf("gRPC error message: %s\n", s.Message())
			if s.Code() == codes.Unavailable {
				t.Log("The service is unavailable. Ensure that the server is running and accessible.")
			}
			t.SkipNow()
		} else {
			log.Fatalf("RPC call failed: %v", err)
		}
	}

	// process response
	idx := 0
	for {
		idx++
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not receive response: %v", err)
			break
		}
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			t.Parallel()
			// Parse the DOT data into the graph
			g := dotGraph.NewDotMultiGraph()
			err = dot.UnmarshalMulti([]byte(resp.GetResult()), g)
			if err != nil {
				log.Fatalf("failed to unmarshal DOT data: %v", err)
			}
			cycles := topo.DirectedCyclesIn(g)

			if len(cycles) == 0 {
				t.Logf("Analysis result: %s\n", resp.GetResult()) // todo: handle command not found
				t.Logf("Cycles: %s", cycles)

				// // Print out the nodes and edges
				// t.Log("Nodes:")
				// for nodes := g.Nodes(); nodes.Next(); {
				// 	node := nodes.Node()
				// 	t.Logf("Node ID: %v\n", node.ID())
				// }

				// t.Log("Edges:")
				// for edges := g.Edges(); edges.Next(); {
				// 	edge := edges.Edge()
				// 	t.Logf("Edge from %v to %v\n", edge.From().ID(), edge.To().ID())
				// }
				// data, err := dot.Marshal(g, "Example Graph", "", "\t")
				// if err != nil {
				// 	log.Fatalf("error marshalling graph to DOT: %v", err)
				// }
				// t.Log(string(data))
				t.Errorf("Expected cycles to be > 0, found none")
			}
		})
	}
}

func TestAlloyAcyclicGraphs(t *testing.T) {
	filePath := "/app/testing/alloy/directed_graph.als"
	alloyCommand := "run acyclic for 4"

	client, conn, err := rpc.NewClient("172.17.0.1:8080")
	if err != nil {
		t.Fatalf("could not create connection client to server: %v", err)
	}
	defer conn.Close()

	stream, err := rpc.UploadFile(client, filePath, alloyCommand)
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			t.Logf("gRPC error code: %s\n", s.Code())
			t.Logf("gRPC error message: %s\n", s.Message())
			if s.Code() == codes.Unavailable {
				t.Log("The service is unavailable. Ensure that the server is running and accessible.")
			}
			t.SkipNow()
		} else {
			log.Fatalf("RPC call failed: %v", err)
		}
	}

	// process response
	idx := 0
	for {
		idx++
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not receive response: %v", err)
			break
		}
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			t.Parallel()
			// Parse the DOT data into the graph
			g := dotGraph.NewDotMultiGraph()
			err = dot.UnmarshalMulti([]byte(resp.GetResult()), g)
			if err != nil {
				log.Fatalf("failed to unmarshal DOT data: %v", err)
			}
			cycles := topo.DirectedCyclesIn(g)

			if len(cycles) > 0 {
				t.Logf("Analysis result: %s\n", resp.GetResult()) // todo: handle command not found

				// // Print out the nodes and edges
				// t.Log("Nodes:")
				// for nodes := g.Nodes(); nodes.Next(); {
				// 	node := nodes.Node()
				// 	t.Logf("Node ID: %v\n", node.ID())
				// }

				// t.Log("Edges:")
				// for edges := g.Edges(); edges.Next(); {
				// 	edge := edges.Edge()
				// 	t.Logf("Edge from %v to %v\n", edge.From().ID(), edge.To().ID())
				// }
				// data, err := dot.Marshal(g, "Example Graph", "", "\t")
				// if err != nil {
				// 	log.Fatalf("error marshalling graph to DOT: %v", err)
				// }
				// t.Log(string(data))

				t.Errorf("Expected cycles to be < 0, found: %s", cycles)
			}
		})
	}
}

func TestParseInputGraph(t *testing.T) {
	var tests = []struct {
		input                 string
		expectedNodeCount     int
		expectedEdgeCount     int
		expectedEdgeTestCases []testEdge
	}{
		{
			// multi directed graph, string node names, with no cycles, comments, and empty lines as adjacency list
			trimWhitespaceFromLines(`# this is a comment
			a b c # this is another comment

			d
			`),
			4,
			2,
			[]testEdge{
				{"a", "a", false},
				{"a", "b", true},
				{"a", "c", true},
				{"a", "d", false},
				{"b", "a", false},
				{"b", "b", false},
				{"b", "c", false},
				{"b", "d", false},
				{"c", "a", false},
				{"c", "b", false},
				{"c", "c", false},
				{"c", "d", false},
				{"d", "a", false},
				{"d", "b", false},
				{"d", "c", false},
				{"d", "d", false},
			},
		},
	}

	for _, test := range tests {
		lg, _ := ParseInputGraph(strings.NewReader(test.input))
		// t.Log(test.input)

		actualNodeCount := lg.Graph.Nodes().Len()
		if actualNodeCount != test.expectedNodeCount {
			t.Errorf(
				"ParseInputGraph(%v).Graph.Nodes().Len() = %v, want %v",
				test.input, actualNodeCount, test.expectedNodeCount,
			)
		}

		actualNumberOfLabels := len(lg.labels)
		if actualNumberOfLabels != test.expectedNodeCount {
			t.Errorf(
				"len(ParseInputGraph(%v).labels) = %v, want %v",
				test.input, actualNumberOfLabels, test.expectedNodeCount,
			)
		}
		if !areValuesDistinct(lg.labels) {
			t.Errorf("areValuesDistinct(%v) = false, want true", lg.labels)
		}

		actualEdgeCount := lg.Graph.Edges().Len()
		if test.expectedEdgeCount != actualEdgeCount {
			t.Errorf(
				"ParseInputGraph(%v).Graph.Edges().Len() = %v, want %v",
				test.input, actualEdgeCount, test.expectedEdgeCount,
			)
		}

		// this may seem excessive but it caught a bug
		expectedTestEdgesCount := test.expectedNodeCount * test.expectedNodeCount
		actualTestEdgesCount := len(test.expectedEdgeTestCases)
		if actualTestEdgesCount != expectedTestEdgesCount {
			t.Errorf("Found %v edge test cases, want %v for full coverage",
				actualTestEdgesCount, expectedTestEdgesCount,
			)
		}

		for _, testEdge := range test.expectedEdgeTestCases {
			actual := lg.Graph.HasEdgeFromTo(lg.labels[testEdge.from], lg.labels[testEdge.to])
			if actual != testEdge.expected {
				t.Errorf(
					"ParseInputGraph(%v).Graph.HasEdgeFromTo(%v, %v) = %v, want %v",
					test.input, testEdge.from, testEdge.to, actual, testEdge.expected,
				)
			}
		}
	}
}

// func TestBuildDirectedGraphOfFile(t *testing.T) {
// 	fmt.Println("TestBuildDirectedGraphFileNotFoundError")
// }

// func TestBuildDirectedGraphOfFileWithStrings(t *testing.T) {
// 	fmt.Println("TestBuildDirectedGraphFileNotFoundError")
// }

// func TestBuildDirectedGraphHandlesEmptyLines(t *testing.T) {
// 	fmt.Println("TestBuildDirectedGraphFileNotFoundError")
// }

func TestRemoveElementsAfterPrefix(t *testing.T) {
	var tests = []struct {
		slice    []string
		prefix   string
		expected []string
	}{
		{[]string{"a", "#b", "c"}, "#", []string{"a"}},
		{[]string{}, "#", []string{}},                    // empty slice
		{[]string{"a"}, "", []string{}},                  // empty prefix
		{[]string{"a#", "#b", "c"}, "#", []string{"a#"}}, // only checks prefix
	}

	for _, test := range tests {
		actual := removeElementsAfterPrefix(test.slice, test.prefix)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf(
				"TestRemoveElementsAfterPrefix(%v, %v) = %v, want %v",
				test.slice, test.prefix, actual, test.expected,
			)
		}
	}
}

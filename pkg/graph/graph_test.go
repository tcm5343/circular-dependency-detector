package graph

import (
	// "regexp"
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"

	pb "github.com/tcm5343/circular-dependency-detector/protos"
	"google.golang.org/grpc"
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

func TestAlloy(t *testing.T) {

	// if len(os.Args) < 2 {
	//     log.Fatalf("Usage: %s <path-to-als-file>", os.Args[0])
	// }
	filePath := os.Args[1]

	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := pb.NewAlloyAnalyzerClient(conn)

	// Define a request
	req := &pb.ModelRequest{
		FilePath: filePath,
	}

	// Call the service
	resp, err := client.AnalyzeModel(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling service: %v", err)
	}

	// Print the response
	fmt.Println("Result from Alloy Analyzer:", resp.Result)
	// Check for the environment variable to decide if the test should run
	if os.Getenv("RUN_MY_FEATURE_TEST") != "true" {
		t.Skip("Skipping TestMyFeature because RUN_MY_FEATURE_TEST is not set to true")
	}

	// Your test logic here
	t.Log("Running TestMyFeature")
	// ... test code ...
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
		fmt.Println(test.input)

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

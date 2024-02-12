package graph

import (
	// "regexp"
	"reflect"
	"strings"
	"testing"
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

func trimWhitespaceFromLines(input string) string { // todo: this should be unit tested too
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimLeft(line, " \t") // trim leading spaces and tabs
	}
	return strings.Join(lines, "\n")
}

func areValuesDistinct(m map[string]int64) bool {
	seen := make(map[int64]struct{})
	for _, v := range m {
		if _, exists := seen[v]; exists {
			return false // If the value has been seen before, not distinct
		}
		seen[v] = struct{}{}
	}
	return true // All values are distinct
}

func TestParseInputGraph(t *testing.T) {
	input := trimWhitespaceFromLines(`
	a b c
    d
	`)
	lg, _ := ParseInputGraph(strings.NewReader(input))

	expectedNodeCount := 4
	actualNodeCount := lg.Graph.Nodes().Len()
	if actualNodeCount != expectedNodeCount {
		t.Errorf(
			"ParseInputGraph(%v).Graph.Nodes().Len() = %v, want %v",
			input, actualNodeCount, expectedNodeCount,
		)
	}

	actualNumberOfLabels := len(lg.labels)
	if actualNumberOfLabels != expectedNodeCount {
		t.Errorf(
			"len(ParseInputGraph(%v).labels) = %v, want %v",
			input, actualNumberOfLabels, expectedNodeCount,
		)
	}
	if !areValuesDistinct(lg.labels) {
		t.Errorf("areValuesDistinct(%v) = false, want true", lg.labels)
	}

	expectedEdgeCount := 2
	actualEdgeCount := lg.Graph.Edges().Len()
	if expectedEdgeCount != actualEdgeCount {
		t.Errorf(
			"ParseInputGraph(%v).Graph.Edges().Len() = %v, want %v",
			input, actualEdgeCount, expectedEdgeCount,
		)
	}

	testEdges := []testEdge{
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
	}

	// this may seem excessive but caught a bug
	expectedTestEdgesCount := expectedNodeCount * expectedNodeCount
	actualTestEdgesCount := len(testEdges)
	if actualTestEdgesCount != expectedTestEdgesCount {
		t.Errorf("Found %v edge test cases, want %v for full coverage",
			actualTestEdgesCount, expectedTestEdgesCount,
		)
	}

	for _, testEdge := range testEdges {
		actual := lg.Graph.HasEdgeFromTo(lg.labels[testEdge.from], lg.labels[testEdge.to])
		if actual != testEdge.expected {
			t.Errorf(
				"ParseInputGraph(%v).Graph.HasEdgeFromTo(%v, %v) = %v, want %v",
				input, testEdge.from, testEdge.to, actual, testEdge.expected,
			)
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

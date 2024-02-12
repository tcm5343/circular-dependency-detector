package graph

import (
	// "regexp"
	// "fmt"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

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

func TestParseInputGraph(t *testing.T) {
	input := strings.TrimSpace(``)
	wg, _ := ParseInputGraph(strings.NewReader(input))
	fmt.Println(wg.Graph.Nodes().Len())
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

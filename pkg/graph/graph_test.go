package graph

import (
	// "regexp"
	// "fmt"
	"reflect"
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

// func TestBuildDirectedGraphOfEmptyFile(t *testing.T) {
// 	fmt.Println("TestBuildDirectedGraphFileNotFoundError")
// }

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
	expected := []string{"a"}
	actual := removeElementsAfterPrefix([]string{"a", "#b", "c"}, "#")
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v != %v", actual, expected)
	}
}

func TestRemoveElementsAfterPrefixOnlyChecksPrefix(t *testing.T) {
	expected := []string{"a#"}
	actual := removeElementsAfterPrefix([]string{"a#", "#b", "c"}, "#")
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v != %v", actual, expected)
	}
}

func TestRemoveElementsAfterPrefixIfEmptySlice(t *testing.T) {
	expected := []string{}
	actual := removeElementsAfterPrefix([]string{}, "#")
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v != %v", actual, expected)
	}
}

func TestRemoveElementsAfterPrefixIfEmptyPrefix(t *testing.T) {
	expected := []string{}
	actual := removeElementsAfterPrefix([]string{"a", "b"}, "")
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v != %v", actual, expected)
	}
}

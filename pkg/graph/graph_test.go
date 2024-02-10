package graph

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestTopologicalGenerationsOf(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg := "Gladys"
	if !want.MatchString(msg) {
		t.Fatalf(`Hello("Gladys") = %q, , want match for %#q, nil`, msg, want)
	}
}

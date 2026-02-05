package coloring_test

import (
	"ascii-art-color/internal/coloring"
	"testing"
)

func TestFindPositions_SingleOccurrence(t *testing.T) {
	text := "hello"
	substring := "e"
	expected := []bool{false, true, false, false, false}
	output := coloring.FindPositions(text, substring)
	if len(output) != len(expected) {
		t.Fatalf("expected length %d,got %d", len(expected), len(output))
	}
	for i := range expected {
		if output[i] != expected[i] {
			t.Errorf(
				"at index %d: expected %v,got %v",
				i,
				expected[i],
				output[i],
			)
		}
	}
}

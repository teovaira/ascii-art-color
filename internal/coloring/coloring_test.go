package coloring_test

import (
	"ascii-art-color/internal/coloring"
	"strings"
	"testing"
)

const (
	green = "\x1b[32m"
	reset = "\x1b[0m"
)

// helper: count how many times color code appears
func countColor(s string) int {
	return strings.Count(s, green)
}

func TestApplyColor_FindPositionsCoverage(t *testing.T) {
	tests := []struct {
		name           string
		text           string
		substring      string
		expectedColors int
	}{
		{
			name:           "single character match",
			text:           "hello",
			substring:      "e",
			expectedColors: 1,
		},
		{
			name:           "single multi-character occurrence",
			text:           "hello",
			substring:      "ll",
			expectedColors: 2,
		},
		{
			name:           "actual multiple occurrences",
			text:           "testtest",
			substring:      "test",
			expectedColors: 8,
		},
		{
			name:           "overlapping occurrences",
			text:           "banana",
			substring:      "ana",
			expectedColors: 5, // b[a n a n a] → positions 1–5
		},
		{
			name:           "substring not found",
			text:           "hello",
			substring:      "z",
			expectedColors: 0,
		},
		{
			name:           "case sensitivity",
			text:           "hello",
			substring:      "H",
			expectedColors: 0,
		},
		{
			name:           "empty substring",
			text:           "abc",
			substring:      "",
			expectedColors: 3,
		},
		{
			name:           "empty text",
			text:           "",
			substring:      "a",
			expectedColors: 0,
		},
		{
			name:           "substring longer than text",
			text:           "hi",
			substring:      "hello",
			expectedColors: 0,
		},
	}

	art := []string{
		"############################",
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			colored := coloring.ApplyColor(
				art,
				tt.text,
				tt.substring,
				green,
			)

			if len(colored) == 0 {
				t.Fatal("ApplyColor returned empty output")
			}

			line := colored[0]
			got := countColor(line)

			if got != tt.expectedColors {
				t.Errorf(
					"expected %d colored characters, got %d",
					tt.expectedColors,
					got,
				)
			}

			// If any coloring happened, reset must also exist
			if got > 0 && !strings.Contains(line, reset) {
				t.Error("expected reset code to be present")
			}
		})
	}
}

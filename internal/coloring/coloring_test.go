package coloring_test

import (
	"ascii-art-color/internal/coloring"
	"strings"
	"testing"
)

const (
	green = "\033[32m"
	reset = "\033[0m"
)

func countColorSegments(s string) int {
	return strings.Count(s, green)
}

func TestApplyColor_FindPositionsCoverage(t *testing.T) {
	tests := []struct {
		name           string
		text           string
		substring      string
		charWidths     []int
		expectedColors int
	}{
		{
			name:           "single character match",
			text:           "hello",
			substring:      "e",
			charWidths:     []int{6, 6, 6, 6, 6},
			expectedColors: 1,
		},
		{
			name:           "single multi-character occurrence",
			text:           "hello",
			substring:      "ll",
			charWidths:     []int{6, 6, 6, 6, 6},
			expectedColors: 1,
		},
		{
			name:           "actual multiple occurrences",
			text:           "testtest",
			substring:      "test",
			charWidths:     []int{6, 6, 6, 6, 6, 6, 6, 6},
			expectedColors: 2,
		},
		{
			name:           "overlapping occurrences",
			text:           "banana",
			substring:      "ana",
			charWidths:     []int{6, 6, 6, 6, 6, 6},
			expectedColors: 2,
		},
		{
			name:           "substring not found",
			text:           "hello",
			substring:      "z",
			charWidths:     []int{6, 6, 6, 6, 6},
			expectedColors: 0,
		},
		{
			name:           "case sensitivity",
			text:           "hello",
			substring:      "H",
			charWidths:     []int{6, 6, 6, 6, 6},
			expectedColors: 0,
		},
		{
			name:           "empty substring",
			text:           "abc",
			substring:      "",
			charWidths:     []int{6, 6, 6},
			expectedColors: 3,
		},
		{
			name:           "empty text",
			text:           "",
			substring:      "a",
			charWidths:     []int{},
			expectedColors: 0,
		},
		{
			name:           "substring longer than text",
			text:           "hi",
			substring:      "hello",
			charWidths:     []int{6, 6},
			expectedColors: 0,
		},
	}

	art := []string{
		"############################################",
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			colored := coloring.ApplyColor(
				art,
				tt.text,
				tt.substring,
				green,
				tt.charWidths,
			)

			if len(colored) == 0 {
				t.Fatal("ApplyColor returned empty output")
			}

			line := colored[0]
			got := countColorSegments(line)

			if got != tt.expectedColors {
				t.Errorf(
					"expected %d color segments, got %d",
					tt.expectedColors,
					got,
				)
			}

			if got > 0 && !strings.Contains(line, reset) {
				t.Error("expected reset code to be present")
			}
		})
	}
}

func TestApplyColor_VariableWidths(t *testing.T) {
	tests := []struct {
		name           string
		asciiArt       []string
		text           string
		substring      string
		charWidths     []int
		expectedColors int
	}{
		{
			name:           "single char colored",
			asciiArt:       []string{"######......"},
			text:           "ab",
			substring:      "a",
			charWidths:     []int{6, 6},
			expectedColors: 1,
		},
		{
			name:           "no chars colored",
			asciiArt:       []string{"######......"},
			text:           "ab",
			substring:      "z",
			charWidths:     []int{6, 6},
			expectedColors: 0,
		},
		{
			name:           "all chars colored",
			asciiArt:       []string{"######......"},
			text:           "ab",
			substring:      "",
			charWidths:     []int{6, 6},
			expectedColors: 2,
		},
		{
			name:           "empty art",
			asciiArt:       []string{},
			text:           "ab",
			substring:      "a",
			charWidths:     []int{6, 6},
			expectedColors: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			colored := coloring.ApplyColor(tt.asciiArt, tt.text, tt.substring, green, tt.charWidths)

			if len(colored) != len(tt.asciiArt) {
				t.Errorf("expected %d lines, got %d", len(tt.asciiArt), len(colored))
			}

			if len(tt.asciiArt) > 0 {
				got := countColorSegments(colored[0])
				if got != tt.expectedColors {
					t.Errorf("expected %d color segments, got %d", tt.expectedColors, got)
				}

				if got > 0 && !strings.Contains(colored[0], reset) {
					t.Error("expected reset code to be present")
				}
			}
		})
	}
}

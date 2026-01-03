package renderer_test

import (
	"ascii-art/renderer"
	"strings"
	"testing"
)

func TestWithRealStandardBanner(t *testing.T) {
	banner, err := parser.LoadBanner("../testdata/standard.txt")
	if err != nil {
		t.Skipf("skipping integration test: %v", err)
	}

	tests := []struct {
		name      string
		input     string
		wantLines int
	}{
		{"simple word", "Hello", 8},
		{"with space", "Hello World", 8},
		{"with numbers", "Hello123", 8},
		{"single newline", "Hello\nWorld", 16},
		{"double newline", "A\n\nB", 17},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := renderer.RendererASCII(tt.input, banner)
			if err != nil {
				t.Fatalf("RendererASCII failed: %v", err)
			}

			lines := strings.Split(output, "\n")
			if len(lines) != tt.wantLines {
				t.Errorf("expected %d lines, got %d", tt.wantLines, len(lines))
			}
		})
	}
}
func TestWithRealShadowBanner(t *testing.T) {
	banner, err := parser.LoadBanner("../testdata/shadow.txt")
	if err != nil {
		t.Skipf("skipping integration test: %v", err)
	}

	output, err := renderer.RendererASCII("A", banner)
	if err != nil {
		t.Fatalf("RendererASCII failed: %v", err)
	}

	lines := strings.Split(output, "\n")
	if len(lines) != 8 {
		t.Errorf("expected 8 lines, got %d", len(lines))
	}
}
func TestWithRealThinkertoyBanner(t *testing.T) {
	banner, err := parser.LoadBanner("../testdata/thinkertoy.txt")
	if err != nil {
		t.Skipf("skipping integration test: %v", err)
	}

	output, err := renderer.RendererASCII("Hello", banner)
	if err != nil {
		t.Fatalf("RendererASCII failed: %v", err)
	}

	lines := strings.Split(output, "\n")
	if len(lines) != 8 {
		t.Errorf("expected 8 lines, got %d", len(lines))
	}
}

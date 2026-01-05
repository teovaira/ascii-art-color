package parser

import (
	"testing"
)

// testing spaces in shadow file
func TestLoadBannerSpaceChar(t *testing.T) {
	banner, err := LoadBanner("../testdata/shadow.txt")
	if err != nil {
		t.Fatalf("LoadBanner failed: %v", err)
	}

	space := ' '
	expected := []string{
		"",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
	}

	actual, ok := banner[space]
	if !ok {
		t.Errorf("banner does not contain space character")
	}

	if len(actual) != len(expected) {
		t.Errorf("expected %d lines for space, got %d", len(expected), len(actual))
	}

	for i, line := range actual {
		if line != expected[i] {
			t.Errorf("line %d: expected %q, got %q", i, expected[i], line)
		}
	}
}

// missing file test
func TestLoadBannerMissingFile(t *testing.T) {
	_, err := LoadBanner("../testdata/nope.txt")
	if err == nil {
		t.Errorf("expected error for missing file, got nil")
	}
}

// testing exclamation
func TestLoadBannerExclamationChar(t *testing.T) {
	banner, err := LoadBanner("../testdata/shadow.txt")
	if err != nil {
		t.Fatalf("LoadBanner Failed: %v", err)
	}
	ex := '!'
	expected := []string{
		"",
		"   ",
		"_| ",
		"_| ",
		"_| ",
		"   ",
		"_| ",
		"   ",
	}
	actual, ok := banner[ex]
	if !ok {
		t.Errorf("banner does not contain '!' character")
	}
	if len(actual) != len(expected) {
		t.Fatalf("expected %d lines for '!', got %d", len(expected), len(actual))
	}
	for i, line := range actual {
		if line != expected[i] {
			t.Errorf("line %d: expected %q, got %q", i, expected[i], line)
		}
	}
}

// testing spaces in standard file
func TestLoadBannerStandardSpace(t *testing.T) {
	banner, err := LoadBanner("../testdata/standard.txt")
	if err != nil {
		t.Fatalf("loadBanner failed: %v", err)
	}
	space := ' '

	actual, ok := banner[space]
	if !ok {
		t.Fatalf("banner does not contain space character")
	}
	if len(actual) != 8 {
		t.Fatalf("expected 8 lines for space, got %d", len(actual))
	}
	expected := []string{
		"",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
	}
	for i, line := range actual {
		if line != expected[i] {
			t.Errorf("line %d: expected %q, got %q", i, expected[i], line)
		}
	}
}

// testing 'A' in shadow file
func TestLoadBannerShadowA(t *testing.T) {
	banner, err := LoadBanner("../testdata/shadow.txt")
	if err != nil {
		t.Fatalf("LoadBanner failed: %v", err)
	}

	ch := 'A'
	actual, ok := banner[ch]
	if !ok {
		t.Fatalf("banner does not contain 'A' character")
	}
	if len(actual) != 8 {
		t.Fatalf("expected 8 lines for 'A', got %d", len(actual))
	}

	expected := []string{
		"",
		"         ",
		"  _|_|   ",
		"_|    _| ",
		"_|_|_|_| ",
		"_|    _| ",
		"_|    _| ",
		"         ",
	}

	for i, line := range actual {
		if line != expected[i] {
			t.Errorf("line %d: expected %q, got %q", i, expected[i], line)
		}
	}
}

func TestLoadBanner_EmptyFile(t *testing.T) {
	_, err := LoadBanner("../testdata/empty.txt")
	if err == nil {
		t.Error("expected error for empty file, got nil")
	}
}

func TestLoadBanner_CorruptedFile(t *testing.T) {
	_, err := LoadBanner("../testdata/corrupted.txt")
	if err == nil {
		t.Error("expected error for corrupted file, got nil")
	}
}

func TestLoadBanner_OversizedFile(t *testing.T) {
	_, err := LoadBanner("../testdata/oversized.txt")
	if err == nil {
		t.Error("expected error for oversized file, got nil")
	}
}

func TestLoadBanner_Thinkertoy(t *testing.T) {
	banner, err := LoadBanner("../testdata/thinkertoy.txt")
	if err != nil {
		t.Fatalf("thinkertoy failed: %v", err)
	}
	if len(banner) != totalChars {
		t.Errorf("expected %d chars, got %d", totalChars, len(banner))
	}
}

func TestLoadBanner_Numbers(t *testing.T) {
	banner, err := LoadBanner("../testdata/standard.txt")
	if err != nil {
		t.Fatalf("LoadBanner failed: %v", err)
	}

	for r := '0'; r <= '9'; r++ {
		lines, ok := banner[r]
		if !ok {
			t.Errorf("missing digit %c", r)
			continue
		}
		if len(lines) != linesPerGlyph {
			t.Errorf("digit %c has %d lines, expected %d",
				r, len(lines), linesPerGlyph)
		}
	}
}
func TestLoadBanner_CompleteCharacterSet(t *testing.T) {
	banner, err := LoadBanner("../testdata/standard.txt")
	if err != nil {
		t.Fatalf("LoadBanner failed: %v", err)
	}

	if len(banner) != totalChars {
		t.Fatalf("expected %d chars, got %d", totalChars, len(banner))
	}

	for r := firstPrintable; r <= lastPrintable; r++ {
		lines, ok := banner[r]
		if !ok {
			t.Errorf("missing char %c (ASCII %d)", r, r)
			continue
		}
		if len(lines) != linesPerGlyph {
			t.Errorf("char %c (ASCII %d) has %d lines, expected %d",
				r, r, len(lines), linesPerGlyph)
		}
	}
}

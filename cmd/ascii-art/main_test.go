// Package main_test contains unit tests for the main package functions.
// Tests verify command-line argument parsing and banner path resolution.
package main

import (
	"testing"
)

// TestParseArgs_NoArguments verifies ParseArgs returns an error when no text argument is provided.
func TestParseArgs_NoArguments(t *testing.T) {
	args := []string{"./ascii-art"}

	_, _, err := ParseArgs(args)

	if err == nil {
		t.Error("Expected error for no arguments, got nil")
	}

	expectedMsg := "usage: go run . \"text\" [banner]"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message: %q, got: %q", expectedMsg, err.Error())
	}
}

// TestParseArgs_TextOnly verifies ParseArgs defaults to "standard" banner when no banner is specified.
func TestParseArgs_TextOnly(t *testing.T) {
	args := []string{"./ascii-art", "Hello"}

	text, banner, err := ParseArgs(args)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if text != "Hello" {
		t.Errorf("Expected text: 'Hello', got: %q", text)
	}

	if banner != "standard" {
		t.Errorf("Expected banner: 'standard', got: %q", banner)
	}
}

// TestParseArgs_TextAndBanner verifies ParseArgs correctly parses both text and banner arguments.
func TestParseArgs_TextAndBanner(t *testing.T) {
	args := []string{"./ascii-art", "Hello", "shadow"}

	text, banner, err := ParseArgs(args)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if text != "Hello" {
		t.Errorf("Expected text: 'Hello', got: %q", text)
	}

	if banner != "shadow" {
		t.Errorf("Expected banner: 'shadow', got: %q", banner)
	}
}

// TestParseArgs_TooManyArguments verifies ParseArgs returns an error when too many arguments are provided.
func TestParseArgs_TooManyArguments(t *testing.T) {
	args := []string{"./ascii-art", "Hello", "shadow", "extra"}

	_, _, err := ParseArgs(args)

	if err == nil {
		t.Error("Expected error for too many arguments, got nil")
	}
}

// TestParseArgs_AllBannerTypes verifies ParseArgs accepts all valid banner type names.
func TestParseArgs_AllBannerTypes(t *testing.T) {
	testCases := []struct {
		args           []string
		expectedBanner string
	}{
		{[]string{"prog", "Hi", "standard"}, "standard"},
		{[]string{"prog", "Hi", "shadow"}, "shadow"},
		{[]string{"prog", "Hi", "thinkertoy"}, "thinkertoy"},
	}

	for _, tc := range testCases {
		_, banner, err := ParseArgs(tc.args)

		if err != nil {
			t.Errorf("Args %v: expected no error, got: %v", tc.args, err)
		}

		if banner != tc.expectedBanner {
			t.Errorf("Args %v: expected banner %q, got: %q",
				tc.args, tc.expectedBanner, banner)
		}
	}
}

// TestParseArgs_EmptyStringText verifies ParseArgs handles empty string as valid text input.
func TestParseArgs_EmptyStringText(t *testing.T) {
	args := []string{"./ascii-art", ""}

	text, banner, err := ParseArgs(args)

	if err != nil {
		t.Errorf("Expected no error for empty string, got: %v", err)
	}

	if text != "" {
		t.Errorf("Expected empty text, got: %q", text)
	}

	if banner != "standard" {
		t.Errorf("Expected banner: 'standard', got: %q", banner)
	}
}

// TestGetBannerPath_ValidBanners verifies GetBannerPath correctly maps banner names to file paths.
func TestGetBannerPath_ValidBanners(t *testing.T) {
	testCases := []struct {
		banner       string
		expectedPath string
	}{
		{"standard", "testdata/standard.txt"},
		{"shadow", "testdata/shadow.txt"},
		{"thinkertoy", "testdata/thinkertoy.txt"},
	}

	for _, tc := range testCases {
		path, err := GetBannerPath(tc.banner)

		if err != nil {
			t.Errorf("Banner %q: expected no error, got: %v", tc.banner, err)
		}

		if path != tc.expectedPath {
			t.Errorf("Banner %q: expected path %q, got: %q",
				tc.banner, tc.expectedPath, path)
		}
	}
}

// TestGetBannerPath_InvalidBanner verifies GetBannerPath returns an error for invalid banner names.
func TestGetBannerPath_InvalidBanner(t *testing.T) {
	banner := "invalid"

	_, err := GetBannerPath(banner)

	if err == nil {
		t.Error("Expected error for invalid banner, got nil")
	}
}

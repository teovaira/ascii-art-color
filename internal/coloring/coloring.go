// Package coloring provides utilities for colorizing
// matching substrings in ASCII art.
//
// The coloring package is responsible for identifying
// substring occurrences in the input text and applying
// ANSI color codes to the corresponding characters
// in the rendered ASCII art.
package coloring

const charWidth = 6

// ApplyColor colorizes all characters in asciiArt that
// correspond to occurrences of substring in text.
func ApplyColor(
	asciiArt []string,
	text string,
	substring string,
	colorCode string,
) []string {
	positions := findPositions(text, substring)

	if len(asciiArt) == 0 || len(positions) == 0 {
		return asciiArt
	}

	result := make([]string, len(asciiArt))
	copy(result, asciiArt)
	reset := "\x1b[0m"

	for i := range asciiArt {
		line := []rune(result[i])
		offset := 0

		for charIdx, shouldColor := range positions {
			if !shouldColor {
				continue
			}

			start := charIdx*charWidth + offset
			end := start + charWidth

			if end > len(line) {
				break
			}

			line = append(line[:start], append([]rune(colorCode), line[start:]...)...)
			offset += len(colorCode)
			end += len(colorCode)

			line = append(line[:end], append([]rune(reset), line[end:]...)...)
			offset += len(reset)
		}

		result[i] = string(line)
	}

	return result
}

// findPositions determines which character positions in a text
// string belong to occurrences of a given substring.
func findPositions(text string, substring string) []bool {
	result := make([]bool, len(text))

	if len(substring) == 0 {
		for i := range result {
			result[i] = true
		}
		return result
	}

	for i := 0; i <= len(text)-len(substring); i++ {
		match := true

		for p := 0; p < len(substring); p++ {
			if text[i+p] != substring[p] {
				match = false
				break
			}
		}

		if match {
			for p := 0; p < len(substring); p++ {
				result[i+p] = true
			}
		}
	}

	return result
}

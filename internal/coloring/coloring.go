// Package coloring provides utilities for colorizing
// matching substrings in ASCII art.
//
// The coloring package is responsible for identifying
// substring occurrences in the input text and applying
// ANSI color codes to the corresponding characters
// in the rendered ASCII art.
package coloring

// ApplyColor colorizes all characters in asciiArt that
// correspond to occurrences of substring in text.
func ApplyColor(
	asciiArt []string,
	text string,
	substring string,
	colorCode string,
	charWidths []int,
) []string {

	if len(asciiArt) == 0 {
		return asciiArt
	}

	positions := findPositions(text, substring)
	reset := "\033[0m"
	result := make([]string, len(asciiArt))

	for i, line := range asciiArt {
		out := ""
		offset := 0
		inColor := false

		for idx := 0; idx < len(positions) && idx < len(charWidths); idx++ {
			if offset >= len(line) {
				break
			}

			end := offset + charWidths[idx]
			if end > len(line) {
				end = len(line)
			}

			start :=
				positions[idx] &&
					(substring == "" ||
						idx == 0 ||
						!positions[idx-1] ||
						(idx+len(substring) <= len(text) &&
							text[idx:idx+len(substring)] == substring))

			if start {
				out += colorCode
				inColor = true
			}

			out += line[offset:end]

			if inColor && (substring == "" ||
				idx+1 >= len(positions) ||
				!positions[idx+1]) {

				out += reset
				inColor = false
			}

			offset = end
		}

		if offset < len(line) {
			out += line[offset:]
		}

		result[i] = out
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

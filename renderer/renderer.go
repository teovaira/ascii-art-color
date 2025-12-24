package renderer

import "strings"

const bannerHeight = 8

func RendererASCII(input string, banner map[rune][]string) string {
	result := ""
	parts := strings.Split(input, "\n")
	// Remove trailing empty string only if input ends with \n
	if len(parts) > 0 && parts[len(parts)-1] == "" {
		parts = parts[0 : len(parts)-1]
	}
	// Handle special case: input is just empty or just "\n"
	if len(parts) == 0 || len(parts) == 1 && parts[0] == "" {
		return result
	}
	for p, line := range parts {
		// Handle empty lines(from consecutive \n\n)
		if line == "" {
			result += "\n"
			continue
		}
		// Render each line of the ASCII art
		for i := 0; i < bannerHeight; i++ {
			for _, ch := range line {
				rows := banner[ch]

				result += rows[i]

			}
			result += "\n"
		}
		// Don't add extra newline after the last part
		if p == len(parts)-1 {
			// Remove the last newline character
			result = result[:len(result)-1]
		}
	}
	return result
}

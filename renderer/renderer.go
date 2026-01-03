package renderer

import (
	"fmt"
	"strings"
)

const bannerHeight = 8

func RendererASCII(input string, banner map[rune][]string) (string, error) {
	var result strings.Builder
	for _, ch := range input {
		if ch == '\n' {
			continue
		}
		if ch < 32 || ch > 126 {
			return "", fmt.Errorf("not printable characters")
		}
	}
	parts := strings.Split(input, "\n")
	// Remove trailing empty string only if input ends with \n
	if len(parts) > 0 && parts[len(parts)-1] == "" {
		parts = parts[0 : len(parts)-1]
	}
	// Handle special case: input is just empty or just "\n"
	if len(parts) == 0 || len(parts) == 1 && parts[0] == "" {
		return "", nil
	}
	if len(banner) == 0 {
		return "", fmt.Errorf("banner is empty")
	}

	for _, line := range parts {
		// Handle empty lines(from consecutive \n\n)
		if line == "" {
			result.WriteString("\n")
			continue
		}

		// Render each line of the ASCII art
		for i := 0; i < bannerHeight; i++ {
			for _, ch := range line {

				value, err := characterValidation(ch, banner)
				if err != nil {
					return "", err
				}
				result.WriteString(value[i])
			}
			result.WriteString("\n")
		}

	}
	output := result.String()
	// Don't add extra newline after the last part
	if output != "" && output[:len(output)-1] == "\n" {
		// Remove the last newline character
		output = output[:len(output)-1]
	}
	return output, nil
}
func characterValidation(ch rune, banner map[rune][]string) ([]string, error) {

	value, exists := banner[ch]
	if exists == false {
		return []string{}, fmt.Errorf("the character does not exist in the banner")
	}
	if len(value) != bannerHeight {
		return []string{}, fmt.Errorf("The character does not have correct number of rows")
	}
	return value, nil
}

package coloring

const charWidth = 6

func ApplyColor(asciiArt []string, positions []bool, colorCode string) []string {
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

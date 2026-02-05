package coloring

func FindPositions(text string, substring string) []bool {
	result := make([]bool, len(text))

	for i := 0; i <= len(text)-1; i++ {
		if text[i] == substring[0] {
			result[i] = true
			break
		}
	}
	return result
}

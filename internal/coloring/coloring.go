package coloring

func FindPositions(text string, substring string) []bool {
	result := make([]bool, len(text))
	if len(substring) == 0 {
		for i := range result {

			result[i] = true
		}

	}
	if len(substring) != 0 {
		for i := 0; i <= len(text)-len(substring); i++ {

			match := true // assume match

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
	}

	return result
}

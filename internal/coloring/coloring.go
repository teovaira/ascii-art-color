package coloring

func FindPositions(text string, substring string) []bool {
	result := make([]bool, len(text))

	// if substring is empty, nothing to do (we'll handle later)
	if len(substring) == 0 {
		return result
	}

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

	return result
}

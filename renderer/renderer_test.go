package renderer

import (
	"testing"
)

func TestEmptyInput(t *testing.T) {
	input := ""
	output := rendererASCII(input)
	if input != output {
		t.Errorf("Error")
	}
}

package flagparser

import (
	"errors"
	"strings"
)

func ParseArgs(args []string) error {
	if len(args) < 2 {
		return errors.New("error")
	}
	if len(args) > 5 {
		return errors.New("error")
	}

	firstTwoLetters := strings.HasPrefix(args[1], "--")
	if firstTwoLetters == false {
		return errors.New("error")
	}
	return nil
}

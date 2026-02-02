package flagparser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	minimumArgs = 2
	maximumArgs = 5
)

var errUsage = errors.New("Usage: go run . [OPTION] [STRING]")

func ParseArgs(args []string) error {
	count := 0
	if len(args) < minimumArgs {
		return errUsage
	}
	if len(args) > maximumArgs {
		return errUsage
	}
	if err := validateColorFlag(args); err != nil {
		return errUsage
	}
	for i, arg := range args {
		if strings.HasPrefix(arg, "--color=") {
			count++
			if i != 1 {
				return errUsage
			}
			if count > 1 {
				return errUsage
			}
		}
	}
	if strings.HasPrefix(args[1], "--color=") && len(args) < 3 {
		return errUsage
	}
	if strings.HasPrefix(args[1], "--color=") {
		checkColorInTheFlag := strings.Split(args[1], "=")
		if len(checkColorInTheFlag) > 1 && checkColorInTheFlag[1] == "" {
			return errUsage
		}
		color := checkColorInTheFlag[1]
		if color != "" {
			if err := validColors(color); err != nil {
				return fmt.Errorf("%v\n%s", err, errUsage)
			}
		}
	}
	return nil

}
func validateColorFlag(args []string) error {
	isItAFlag := strings.HasPrefix(args[1], "-")
	if isItAFlag {

		firstTwoLetters := strings.HasPrefix(args[1], "--")

		if !firstTwoLetters {
			return errUsage
		}

		hasEqual := strings.Contains(args[1], "=")
		if !hasEqual {
			return errUsage
		}

	}
	return nil
}
func validColors(color string) error {
	allowedColors := map[string]bool{
		"red":     true,
		"green":   true,
		"yellow":  true,
		"orange":  true,
		"blue":    true,
		"magenta": true,
	}
	if _, exists := allowedColors[color]; exists {
		return nil
	}

	if strings.HasPrefix(color, "rgb(") {
		inner := strings.TrimSuffix(strings.TrimPrefix(color, "rgb("), ")")
		separatedText := strings.Split(inner, ",")
		if len(separatedText) != 3 {
			return errors.New("invalid RGB format: expected rgb(r,g,b)")
		}
		for i := 0; i < len(separatedText); i++ {
			digits, err := strconv.Atoi(separatedText[i])
			if err != nil {
				return errors.New("invalid RGB value: must be a number")
			}
			if digits < 0 || digits > 255 {

				return errors.New("invalid RGB value: must be between 0 and 255")

			}
		}
		return nil
	}
	colorValue, checkHex := strings.CutPrefix(color, "#")

	if checkHex {
		if colorValue == "" {
			return errors.New("invalid HEX color: missing hexadecimal value")
		}
		if len(colorValue) != 6 {
			return errors.New("invalid HEX color: expected 6 hexadecimal characters")
		}
		for _, ch := range colorValue {
			if !(ch >= '0' && ch <= '9') && !(ch >= 'a' && ch <= 'f') && !(ch >= 'A' && ch <= 'F') {
				return errors.New("invalid HEX color: contains non-hexadecimal character")
			}
		}

		return nil
	}
	return errors.New("unsupported color format")
}

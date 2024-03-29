package utils

import (
	"errors"
	"unicode"
)

func validateInput(str string) error {
	for _, char := range str {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			return errors.New("only enter numbers or english letters")
		}

	}

	if len(str) > 5 {
		return errors.New("only enter 5 characters")
	}

	return nil
}

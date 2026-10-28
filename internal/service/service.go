package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

var argException = errors.New("Incorrect input data")

func IsOnlySpaces(text string) bool {
	var temp = strings.TrimSpace(text)
	return len(temp) == 0
}

func Convert(inputText string) (string, error) {
	var result string

	if IsOnlySpaces(inputText) {
		return "", fmt.Errorf("Input data: %w", argException)
	}

	if temp := strings.ToLower(inputText); temp == inputText {
		result = morse.ToText(inputText)
		return result, nil
	}

	return morse.ToMorse(inputText), nil
}

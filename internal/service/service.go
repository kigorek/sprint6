package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// MorseOrText функция автоматического определения кода Морзе или обычного текста из переданной строки.

func MorseOrText(anyText string) (string, error) {
	if anyText == "" {
		return "", errors.New("empty text")
	}

	textTrim := strings.TrimSpace(anyText)

	if strings.ContainsAny(textTrim, "-") && strings.ContainsAny(textTrim, ".") {
		return morse.ToText(anyText), nil
	}
	return morse.ToMorse(anyText), nil
}

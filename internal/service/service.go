package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Translate(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", fmt.Errorf("input err")
	}

	if isMorse(trimmed) {
		result := morse.ToText(trimmed)
		if result == "" && trimmed != "" {
			return "", fmt.Errorf("conv to text err")
		}
		return result, nil
	}
	result := morse.ToMorse(trimmed)
	if strings.TrimSpace(result) == "" && strings.TrimSpace(trimmed) != "" {
		return "", fmt.Errorf("conv to morse err")
	}
	return result, nil
}
func isMorse(s string) bool {
	for _, ch := range s {
		if ch != ' ' && ch != '.' && ch != '-' {
			return false
		}
	}
	return true
}

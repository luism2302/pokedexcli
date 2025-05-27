package input

import (
	"strings"
)

func CleanInput(text string) []string {
	if text == "" {
		return []string{}
	}

	lower := strings.ToLower(text)
	trimmed := strings.Fields(lower)

	return trimmed
}

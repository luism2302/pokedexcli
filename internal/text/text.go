package text

import (
	"strings"
)

func CleanInput(text string) []string {
	words := strings.Fields(text)
	cleanWords := make([]string, len(words))
	for i, word := range words {
		cleanWords[i] = strings.ToLower(word)
	}

	return cleanWords
}

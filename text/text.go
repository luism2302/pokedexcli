package text

import (
	"strings"
)

func cleanInput(text string) []string {
	words := strings.Fields(text)
	cleanWords := make([]string, len(words))
	for i, word := range words {
		cleanWords[i] = strings.ToLower(word)
	}

	return cleanWords
}

package engine

import (
	"regexp"
	"strings"
)

type token string

func tokenizer(input string) []token {
	output := []token{}
	spacelessInput := strings.Split(input, " ")
	for _, word := range spacelessInput {
		if len(word) > 3 {
			normalized := normalizeInput(word) //
			// normalized := strings.TrimSpace(word) // Swap normalize
			output = append(output, token(normalized))
		}
	}

	return output
}

func normalizeInput(input string) string {
	normalized := strings.TrimSpace(input)
	normalized = strings.ToLower(normalized)
	normalized = regexp.MustCompile("\n|\t|\r|[^a-zAz]").ReplaceAllString(normalized, "")

	sufixs := []string{"ndo", "nda", "ing", "'ll", "n't", "a", "o"}
	for _, sufix := range sufixs {
		unSufixed := strings.TrimSuffix(normalized, sufix)
		if len(unSufixed) >= 2 {
			normalized = unSufixed
		}
	}

	return normalized
}

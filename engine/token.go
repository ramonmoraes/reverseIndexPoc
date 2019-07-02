package engine

import (
	"strings"
)

type token string;

func tokenizer(input string) []token {
	output := []token{}
	spacelessInput := strings.Split(input, " ")
	for _, word := range spacelessInput {
		if len(word) > 3 {
			normalized := normalizeInput(word)
			output = append(output, token(normalized))
		}
	}

	return output
}

func normalizeInput(input string) string {
	normalized := strings.TrimSpace(input)
	normalized = strings.ToLower(normalized)
	
	sufixs := []string{"ndo", "nda", "ing","'ll", "n't"}
	for _, sufix:= range(sufixs) {
		unSufixed := strings.TrimSuffix(normalized, sufix)
		if len(unSufixed) >=2 {
			normalized = unSufixed
		}
	}

	return normalized
}
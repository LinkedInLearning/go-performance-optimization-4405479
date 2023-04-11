package tokenizer

import (
	"regexp"
	"strings"
)

var (
	wordRe   = regexp.MustCompile(`[a-zA-Z]+`)
	suffixes = []string{"ed", "ing", "s"}
)

func Stem(word string) string {
	for _, s := range suffixes {
		if strings.HasSuffix(word, s) {
			return word[:len(word)-len(s)]
		}
	}

	return word
}

func initialSplit(text string) []string {
	return wordRe.FindAllString(text, -1)
}

func Tokenize(text string) []string {
	var tokens []string
	for _, tok := range initialSplit(text) {
		tok = strings.ToLower(tok)
		tok = Stem(tok)
		tokens = append(tokens, tok)
	}
	return tokens
}

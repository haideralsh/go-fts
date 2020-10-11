package main

import (
	"strings"
	"unicode"
)

func tokenize(text string) []string {
	// takes a string and split it according to the rules inside the anonymous
	// function
	return strings.FieldsFunc(text, func(r rune) bool {
		// Split on any character that is not a letter or a number.
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

// Make everything lowercase
func lowercaseFilter(tokens []string) []string {
	// new slice of strings using the length of all tokens
	r := make([]string, len(tokens))

	// take each token and make it lowercase and send it to the save it in the
	// temp array
	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}
	return r
}

// Remove stop words (commonly used words)
// basically a map of just keys, i.e a set. Since Go doesn't have built-in sets.
var stopwords = map[string]struct{}{
	"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
	"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
}

func stopwordFilter(tokens []string) []string {
	// Same thing, init an empty slice with the length of the tokens
	r := make([]string, 0, len(tokens))

	for _, token := range tokens {
		// in Go when you do myMap["key"] it returns two things
		// the value for that key, and a boolean whether it was
		// found or not.
		if _, ok := stopwords[token]; !ok {
			r = append(r, token)
		}
		// can also be written as:
		//_, found := stopwords[token]
		//if !found {
		// r = append(r, token)
		//}
	}
	return r
}

func process(text string) []string {
	tokens := tokenize(text)
	tokens = lowercaseFilter(tokens)
	tokens = stopwordFilter(tokens)

	return tokens
}

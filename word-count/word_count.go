// Package wordcount provides utilities for counting word frequencies in text.
// It handles Unicode text correctly and treats contractions as single words.
package wordcount

import (
	"strings"
	"unicode"
)

// Frequency maps each word with its count in the phrase.
type Frequency map[string]int

const quote = "'"

// wordSplitter determines if a rune should be considered a word boundary.
// It splits on spaces and punctuation, but preserves apostrophes in contractions.
func wordSplitter(r rune) bool {
	return unicode.IsSpace(r) || (unicode.IsPunct(r) && r != '\'') || unicode.IsSymbol(r)
}

// isStandalonePunct checks if a word is a single punctuation mark.
func isStandalonePunct(word string) bool {
	return len(word) == 1 && unicode.IsPunct(rune(word[0]))
}

// WordCount returns a frequency map of all words in the provided phrase.
// Words are converted to lowercase, and punctuation is handled appropriately.
func WordCount(phrase string) Frequency {
	// Early return if phrase is empty
	if phrase == "" {
		return make(Frequency)
	}

	// Convert to lowercase and split
	words := strings.FieldsFunc(strings.ToLower(phrase), wordSplitter)

	// Initialize the frequency map with rough estimated capacity
	freq := make(Frequency, len(words)/2)

	// Count occurrences
	for _, word := range words {
		// Skip empty strings and standalone punctuation
		if word == "" || isStandalonePunct(word) {
			continue
		}

		// Trim leading/trailing single quotes to handle quoted words consistently
		cleaned := strings.Trim(word, quote)
		if cleaned != "" {
			freq[cleaned]++
		}
	}
	return freq
}

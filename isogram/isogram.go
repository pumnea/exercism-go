// Package isogram contains utilities to check if a word is an isogram.
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram return true if input is an isogram.
// Isogram is a word or phrase without a repeating letter,
// however spaces and hyphens are allowed to appear multiple times.
func IsIsogram(word string) bool {
	if len(word) == 0 {
		return true
	}
	word = strings.ToLower(word)
	seenLetters := make(map[rune]bool)
	for _, r := range word {
		if !unicode.IsLetter(r) {
			continue
		}
		if _, exists := seenLetters[r]; exists {
			return false
		}
		seenLetters[r] = true
	}
	return true
}

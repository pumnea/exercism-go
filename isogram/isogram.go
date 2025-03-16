// Package isogram contains utilities to check if a word is an isogram.
package isogram

import "strings"

// IsIsogram return true if input is an isogram.
// Isogram is a word or phrase without a repeating letter,
// however spaces and hyphens are allowed to appear multiple times.
func IsIsogram(word string) bool {
	if len(word) == 0 {
		return true
	}
	letters := make(map[rune]bool)
	for _, r := range strings.ToLower(word) {
		if r != ' ' && r != '-' {
			_, exists := letters[r]
			if exists {
				return false
			}
			letters[r] = true
		}
	}
	return true
}

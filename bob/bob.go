// Package bob provides function that returns various strings based on input
package bob

import (
	"strings"
	"unicode"
)

// Hey returns a string based on input
// "Sure." if input ends with a question mark.
// "Whoa, chill out!" if input is ALL CAPITAL LETTERS.
// "Calm down, I knoww what I'm doing!" if input is in CAPITAL LETTERS and ends in question mark.
// "Fine. Be that way!" if input is empty or various combinations of whitespace characters.
// "Whatever." for anything else.
func Hey(remark string) string {
	normalizedInput := strings.TrimSpace(remark)

	// Check if it's a question, ends in ?
	isQuestion := strings.HasSuffix(normalizedInput, "?")

	// Check if all letters are uppercase and contains at least one letters
	isYelling := isAllUpper(normalizedInput) && hasLetters(normalizedInput)

	switch {
	case normalizedInput == "":
		return "Fine. Be that way!"
	case isYelling && isQuestion:
		return "Calm down, I know what I'm doing!"
	case isYelling:
		return "Whoa, chill out!"
	case isQuestion:
		return "Sure."
	default:
		return "Whatever."
	}
}

func hasLetters(normalizedInput string) bool {
	for _, r := range normalizedInput {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func isAllUpper(normalizedInput string) bool {
	// If no letters, can't be yelling
	if !hasLetters(normalizedInput) {
		return false
	}
	// All letters must be uppercase
	for _, r := range normalizedInput {
		if unicode.IsLetter(r) && !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

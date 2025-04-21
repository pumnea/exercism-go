// Package acronym provides utilities to create acronyms from strings.
// For example, Abbreviate("As Soon As Possible") returns "ASAP".
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate returns the acronym formed by the first letter of each word in s.
// Words are split on spaces or hyphens; non-letter words are skipped.
// Leading non-letters (e.g., apostrophes in "It's") are trimmed from each word using TrimLeftFunc.
// Example:
//
//	Abbreviate("Liquid-crystal display") // Returns "LCD"
//	Abbreviate("It's-a-test") // Returns "IAT"
func Abbreviate(s string) string {
	var result strings.Builder
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == '-'
	})
	for _, word := range words {
		trimmed := strings.TrimLeftFunc(word, func(r rune) bool {
			return !unicode.IsLetter(r)
		})
		if trimmed != "" {
			result.WriteRune(unicode.ToUpper(rune(trimmed[0])))
		}
	}
	return result.String()
}

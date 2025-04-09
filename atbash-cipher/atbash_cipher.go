// Package atbash provides utilities to encode a string using atbash cipher.
package atbash

import (
	"strings"
	"unicode"
)

// Atbash encrypts input using atbash cipher.
// Atbash cipher relies on transposing all the letters in the alphabet such that the resulting alphabet is backwards.
// The first letter is replaced with the last letter, the second with the second-last, and so on.
// Plain:  abcdefghijklmnopqrstuvwxyz
// Cipher: zyxwvutsrqponmlkjihgfedcba
func Atbash(s string) string {
	var result strings.Builder
	input := strings.ToLower(s)
	count := 0

	for _, r := range input {
		switch {
		case unicode.IsLetter(r):
			if count == 5 {
				result.WriteRune(' ')
				count = 0
			}
			// Encode the letter
			result.WriteRune('a' + (25 - (r - 'a')))
			count++
		case unicode.IsDigit(r):
			result.WriteRune(r)
			count++
		}
	}

	return result.String()
}

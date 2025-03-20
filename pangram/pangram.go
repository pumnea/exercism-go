// Package pangram provides utilities to check for pangrams in strings.
package pangram

import (
	"strings"
)

// IsPangram checks if the input string contains all 26 letters of the English alphabet (a-z),
// ignoring case and non-letter characters. Returns true if it’s a pangram, false otherwise.
func IsPangram(input string) bool {
	// Early exit if input has length less then letters in alphabet.
	if len(input) < 26 {
		return false
	}
	// Array to keep track of seen letters.
	seen := [26]bool{}
	count := 0
	// Convert to lowercase.
	for _, r := range strings.ToLower(input) {
		// Check if r is a letter.
		if r >= 'a' && r <= 'z' {
			// Rune is alias for int32, 'a' -> 97 and 'z' -> 122
			// Index is constructing based int32 value of 'r' - 97.
			index := r - 'a'
			if !seen[index] {
				seen[index] = true
				count++
				// Early exit if all lettters are found.
				if count == 26 {
					return true
				}
			}
		}
	}
	return count == 26
}

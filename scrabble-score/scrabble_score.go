// Package scrabble provides utilities to calculate Scrabble scores for words.
package scrabble

import (
	"unicode"
)

// Score returns the Scrabble score of a word based on its letter values.
func Score(word string) int {
	if len(word) == 0 {
		return 0
	}
	score := 0
	for _, l := range word {
		switch unicode.ToUpper(l) {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			score += 1
		case 'D', 'G':
			score += 2
		case 'B', 'C', 'M', 'P':
			score += 3
		case 'F', 'H', 'V', 'W', 'Y':
			score += 4
		case 'K':
			score += 5
		case 'J', 'X':
			score += 8
		case 'Q', 'Z':
			score += 10
		}
	}
	return score
}

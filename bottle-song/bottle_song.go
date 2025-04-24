// Package bottlesong generates lyrics for the "Ten Green Bottles" song.
package bottlesong

import (
	"fmt"
	"strings"
)

// numberWords maps integers to their English word representations.
var numberWords = map[int]string{
	0:  "no",
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
}

// Recite generates the lyrics for the specified number of verses starting from startBottles.
// It returns an error if the inputs are invalid (e.g., negative bottles or verses).
func Recite(startBottles, verses int) []string {
	var result []string
	for i := 0; i < verses; i++ {
		bottles := startBottles - i
		verseLines := generateVerse(bottles)
		result = append(result, verseLines...)
		if i < verses-1 {
			result = append(result, "") // Add blank line between verses, except after the last one.
		}
	}
	return result
}

// generateVerse creates the lines for a single verse with the given number of bottles.
func generateVerse(bottles int) []string {
	return []string{
		formatBottleLine(bottles, true) + " hanging on the wall,",
		formatBottleLine(bottles, true) + " hanging on the wall,",
		"And if one green bottle should accidentally fall,",
		"There'll be " + formatBottleLine(bottles-1, false) + " hanging on the wall.",
	}
}

// formatBottleLine formats a line describing the number of bottles, optionally capitalizing the number word.
func formatBottleLine(bottles int, capitalize bool) string {
	var b strings.Builder
	word := numberToWord(bottles)
	if capitalize {
		word = strings.ToUpper(word[:1]) + word[1:]
	}
	b.WriteString(word)
	b.WriteString(" green bottle")
	if bottles != 1 {
		b.WriteString("s")
	}
	return b.String()
}

// numberToWord converts a number to its word representation, defaulting to numeric string for unknown numbers.
func numberToWord(n int) string {
	if word, ok := numberWords[n]; ok {
		return word
	}
	return fmt.Sprint(n)
}

// Package anagram provides utilities to check for anagrams
package anagram

import (
	"sort"
	"strings"
)

// Detect finds all anagrams of subject within candidates.
// An anagram is a word formed by rearranging the letters of another word,
// excluding the original word itself (case-insensitive).
func Detect(subject string, candidates []string) []string {
	// Slice of anagrams for subject
	anagrams := make([]string, 0, len(candidates))

	// Make sorted and lowercased slice from subject
	slSubj := lowerAndSort(subject)

	for _, candidate := range candidates {
		// Omit candidates that are equal to subject and with lenght not equal to subject
		if len(candidate) == len(subject) && !strings.EqualFold(subject, candidate) {

			// Make sorted and lowercased slice from candidate
			slCand := lowerAndSort(candidate)

			// Compare slices converted to strings
			if string(slSubj) == string(slCand) {
				anagrams = append(anagrams, candidate)
			}

		}
	}

	return anagrams
}

// lowerAndSort converts a string to lowercase and returns its sorted runes.
// This is a helper function for anagram detection.
func lowerAndSort(input string) []rune {
	runes := []rune(strings.ToLower(input))
	sort.Slice(runes, func(i int, j int) bool {
		return runes[i] < runes[j]
	})
	return runes
}

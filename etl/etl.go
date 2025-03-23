// Package etl provides functions for transforming data structures, extract-transform-load (ETL) processes.
package etl

import (
	"strings"
)

// Transform returns a map of score to map of letters into a map of lowercase letters to score
// For example, given {1: ["A", "E"]}, it returns {"a": 1, "e": 1}
func Transform(in map[int][]string) map[string]int {
	result := make(map[string]int)
	for score, letters := range in {
		for _, letter := range letters {
			result[strings.ToLower(letter)] = score
		}
	}
	return result
}

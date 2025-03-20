// Package reverse provides utilities to reverse a string.
package reverse

import (
	"strings"
)

// Reverse returns a new string that is reverse of the input.
func Reverse(input string) string {
	// Early return for an empty string.
	if len(input) == 0 {
		return ""
	}

	// Convert input to []runes.
	runes := []rune(input)
	length := len(runes)

	// Initialize builder with pre-alocated capacity.
	var builder strings.Builder
	builder.Grow(length)

	for i := length - 1; i >= 0; i-- {
		builder.WriteRune(runes[i])
	}

	return builder.String()
}

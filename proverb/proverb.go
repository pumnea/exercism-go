// Package proverb generates a proverbial rhyme from a list of inputs.
package proverb

import "strings"

// Proverb generates a proverb from a list of items, following the "horseshoe nail" pattern.
// For each pair of consecutive items, it creates a line "For want of a X the Y was lost."
// The final line is "And all for the want of a Z," where Z is the first item.
// Returns an empty slice if the input is empty.
func Proverb(rhyme []string) []string {
	// Return empty slice for empty input
	if len(rhyme) == 0 {
		return []string{}
	}

	// Initialize result slice with capacity for efficiency
	result := make([]string, 0, len(rhyme)+1)

	// Generate lines for each pair of consecutive items
	for i := 0; i < len(rhyme)-1; i++ {
		line := strings.Builder{}
		line.WriteString("For want of a ")
		line.WriteString(rhyme[i])
		line.WriteString(" the ")
		line.WriteString(rhyme[i+1])
		line.WriteString(" was lost.")
		result = append(result, line.String())
	}

	// Add the final line referencing the first item, if there’s at least one item
	if len(rhyme) > 0 {
		final := strings.Builder{}
		final.WriteString("And all for the want of a ")
		final.WriteString(rhyme[0])
		final.WriteString(".")
		result = append(result, final.String())
	}

	return result
}

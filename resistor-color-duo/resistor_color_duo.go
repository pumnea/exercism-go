// Package resistorcolorduo provides utilities to calculate resistance based on two color bands.
package resistorcolorduo

// Value returns the resistance value of a resistor based on its first two color bands.
// The colors are mapped to digits (e.g., "black" = 0, "brown" = 1), concatenated, and converted to an integer.
// Returns 0 if fewer than 2 colors are provided or if the colors are invalid.
func Value(colors []string) int {
	// Return 0 if fewer than 2 colors are provided
	if len(colors) < 2 {
		return 0
	}
	// Map of color names to their digit values
	values := map[string]int{
		"black":  0,
		"brown":  1,
		"red":    2,
		"orange": 3,
		"yellow": 4,
		"green":  5,
		"blue":   6,
		"violet": 7,
		"grey":   8,
		"white":  9,
	}
	// Build the resistance value from the first two colors
	total := 0
	// Multiply by 10 to get the expected result 10 + 5 = 15 and not 1 + 5 = 6
	if value, ok := values[colors[0]]; ok {
		total += value * 10
	}
	if value, ok := values[colors[1]]; ok {
		total += value
	}

	return total
}

// Package raindrops converts numbers to strings based on their factors.
package raindrops

import "strconv"

// Convert returns a string of raindrop sounds for factors 3, 5, and 7, or the number if none apply.
func Convert(number int) string {
	switch {
	case number%3 == 0 && number%5 == 0 && number%7 == 0:
		return "PlingPlangPlong"
	case number%3 == 0 && number%5 == 0:
		return "PlingPlang"
	case number%3 == 0 && number%7 == 0:
		return "PlingPlong"
	case number%5 == 0 && number%7 == 0:
		return "PlangPlong"
	case number%3 == 0:
		return "Pling"
	case number%5 == 0:
		return "Plang"
	case number%7 == 0:
		return "Plong"
	default:
		return strconv.Itoa(number)
	}
}

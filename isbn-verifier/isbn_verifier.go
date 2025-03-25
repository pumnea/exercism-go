// Package isbn provides utility to check validity per isbn10 standard.
package isbn

import "strings"

// IsValidISBN returns true if input is a valid isbn10.
func IsValidISBN(isbn string) bool {
	// Remove leading/trailing whitespaces and hyphens
	sanitizedInput := strings.ReplaceAll(strings.TrimSpace(isbn), "-", "")

	// Check lenght (ISBN10 should be 10 characters long)
	if len(sanitizedInput) != 10 {
		return false
	}
	// Init slice with size of sanitizedInput
	digits := make([]int, len(sanitizedInput))

	// Parse each character into an integer
	for i, r := range sanitizedInput {
		switch r {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			digits[i] = int(r - '0')
		case 'X':
			if i == 9 {
				digits[i] = 10
			} else {
				return false
			}
		default:
			return false
		}
	}

	sum := 0

	// Calculate sum of digits multiplied by weight (from 10 to 1)
	for i, v := range digits {
		sum += v * (10 - i)
	}

	return sum%11 == 0
}

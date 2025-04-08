// Package lsproduct provides utilities for calculating the largest product of
// contiguous digit sequences in a string, handling Unicode digits and
// preserving mathematical correctness for edge cases.
package lsproduct

import (
	"errors"
	"unicode"
)

// LargestSeriesProduct calculates the largest product of any contiguous sequence
// of digits in the input string with the given span length.
//
// The digits parameter must contain only numeric characters (0-9). The span must
// be non-negative and not exceed the length of digits. If span is 0, returns 1
// (the empty product). Returns an error for invalid inputs.
//
// Examples:
//   - "12345", span 3 → 60 (from "345")
//   - "29", span 2 → 18
//   - "", span 0 → 1
func LargestSeriesProduct(digits string, span int) (int64, error) {
	// Validate span
	if span < 0 {
		return 0, errors.New("span cannot be negative")
	}

	if span > len(digits) {
		return 0, errors.New("span must be smaller than string length")
	}

	// Handle special case for empty product
	if span == 0 {
		return 1, nil
	}

	// Check for insufficient digits
	if len(digits) == 0 {
		return 0, errors.New("empty string requires span of 0")
	}

	// Check for non-digit characters
	for _, r := range digits {
		if !unicode.IsDigit(r) {
			return 0, errors.New("digits input must only contain digits")
		}
	}

	var maxProduct int64

	// Slide window across digits, computing products
	for i := 0; i <= len(digits)-span; i++ {
		product := int64(1)

		// Calculate product for current window
		for j := i; j < i+span; j++ {
			digit := int64(digits[j] - '0')

			// Skip further multiplication if we hit a zero
			if digit == 0 {
				product = 0
				break
			}

			product *= digit
		}

		// Update maximum product (first window or new maximum)
		if i == 0 || product > maxProduct {
			maxProduct = product
		}
	}

	return maxProduct, nil
}

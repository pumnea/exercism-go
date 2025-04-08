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
	// Validate span and handle special case
	if err := validateSpan(span, len(digits)); err != nil {
		return 0, err
	}

	// Special case for span=0
	if span == 0 {
		return 1, nil
	}

	// Validate digits
	if err := validateDigits(digits); err != nil {
		return 0, err
	}

	var maxProduct int64

	// Slide window across digits, computing products
	for i := 0; i <= len(digits)-span; i++ {
		// Extract the current window substring
		window := digits[i : i+span]
		product := int64(1)

		// Calculate product for current window
		product = calcWindowProduct(window, product)

		// Update maximum product (first window or new maximum)
		if i == 0 || product > maxProduct {
			maxProduct = product
		}
	}

	return maxProduct, nil
}

// validateDigits checks if digits is greater than 0 and it contains only digits.
func validateDigits(digits string) error {
	// Check for insufficient digits
	if len(digits) == 0 {
		return errors.New("empty string requires span of 0")
	}

	// Check for non-digit characters
	for _, r := range digits {
		if !unicode.IsDigit(r) {
			return errors.New("digits input must only contain digits")
		}
	}
	return nil
}

// validateSpan checks if the span is valid for the given string length.
func validateSpan(span, length int) error {
	if span < 0 {
		return errors.New("span cannot be negative")
	}
	if span > length {
		return errors.New("span must not exceed string length")
	}
	if span > 0 && length == 0 {
		return errors.New("empty string requires span of 0")
	}
	return nil
}

// calcWindowProduct returns product of the current window.
func calcWindowProduct(window string, product int64) int64 {
	for _, r := range window {
		digit := int64(r - '0')

		// Skip further multiplication if we hit a zero
		if digit == 0 {
			product = 0
			break
		}

		product *= digit
	}
	return product
}

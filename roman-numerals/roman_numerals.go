// Package romannumerals provides functionality to convert Arabic numerals to traditional Roman numerals.
// It supports numbers from 1 to 3999, following standard Roman numeral conventions (e.g., subtractive notation like IV for 4).
// The implementation uses a greedy algorithm for efficient conversion.
package romannumerals

import "errors"

// ToRomanNumeral converts an Arabic numeral to a Roman numeral string.
// Returns an error if the number is not in the range 1 to 3999.
func ToRomanNumeral(input int) (string, error) {
	// Early error return if number is not within the range 1 .. 3999.
	if input < 1 || input > 3999 {
		return "", errors.New("number must be between 1 and 3999")
	}
	// Roman numeral values and their symbols in descinding order.
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""
	remaing := input

	for i, v := range values {
		for remaing >= v {
			result += symbols[i]
			remaing -= v
		}
	}

	return result, nil
}

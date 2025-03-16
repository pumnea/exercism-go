// Package luhn checks if input is a valid Luhn formula.
// Luhn formula:
// step 1: double every second digit, starting from the right.
// if doubling the number results in a number greater than 9 then subtract 9 from the product.
// step 2: sum all of the digits.
// step 3: check if sum is evenly divisible by 10, then the number is valid.
package luhn

import (
	"strings"
	"unicode"
)

// Valid returns true if input is valid.
func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}
	sum := 0
	for i := len(id) - 1; i >= 0; i-- {
		r := rune(id[i])
		if !unicode.IsNumber(r) {
			return false
		}
		num := int(r - '0')
		if (len(id)-i-1)%2 == 1 {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}
		sum += num
	}
	return sum%10 == 0
}

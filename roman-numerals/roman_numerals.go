package romannumerals

import "errors"

// M	    D	  C	  L	  X	  V	I
// 1000	500	100	50	10	5	1
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

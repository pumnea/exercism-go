// Package armstrong provides utilities for checking Armstrong numbers.
// An Armstrong number is a non-negative integer that equals the sum of its digits,
// each raised to the power of the number of digits.
package armstrong

// IsNumber returns true if 'n' is an armstrong number.
// An Armstrong number is a number that is the sum of
// its own digits each raised to the power of the number of digits.
func IsNumber(n int) bool {
	if n < 0 {
		return false
	}

	if n == 0 {
		return true
	}
	digits := split(n)
	return n == sum(digits)
}

// pow computes base^exp using integer arithmetic.
func pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

// sum returns sum of the digits raised to power of length of 'num'
func sum(nums []int) int {
	sum := 0
	for _, d := range nums {
		sum += pow(d, len(nums))
	}
	return sum
}

// split returns a sorted []int of digits of 'n'
func split(n int) []int {
	result := make([]int, 0)
	for n > 0 {
		digit := n % 10
		result = append(result, digit)
		n /= 10
	}
	return result
}

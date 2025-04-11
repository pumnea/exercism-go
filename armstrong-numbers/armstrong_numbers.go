// Package armstrong provides utlities to check if a given number is an armstrong number.
package armstrong

import (
	"math"
	"sort"
)

// IsNumber returns true if 'n' is an armstrong number.
// An Armstrong number is a number that is the sum of
// its own digits each raised to the power of the number of digits.
func IsNumber(n int) bool {
	digits := split(n)
	digitsSum := sum(digits)
	return isArm(n, digitsSum)
}

// isArm returns true if 'n' is equal to sum of each digit raised
// to power of number of digits.
func isArm(n, sumN int) bool {
	return n == sumN
}

// sum returns sum of the digits raised to power of length of 'num'
func sum(nums []int) int {
	sum := 0
	for _, d := range nums {
		sum += int(math.Pow(float64(d), float64(len(nums))))
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
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result
}

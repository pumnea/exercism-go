// Package diffsquares provides utilities to obtain difference between
// the square of the sum and the sum of the squares of the first N natural numbers.
package diffsquares

// SquareOfSum() returns square of sum of first N natural numbers.
func SquareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

// SumOfSquares() returns sum of squares of first N natural numbers.
func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

// Difference() returns difference between SquareOfSum and SumOfSquares.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

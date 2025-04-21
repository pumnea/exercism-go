// Package prime provides utilities to compute the prime factors of a natural number.
// A prime factor is a prime number that divides the input evenly, and 1 is not considered prime.
package prime

// Factors computes the prime factors of a given natural number n.
// It returns a slice of prime factors in ascending order.
// If n <= 0, it returns an empty slice.
func Factors(n int64) []int64 {
	if n <= 0 {
		return nil
	}
	var factors []int64
	for divisor := int64(2); divisor*divisor <= n; divisor++ {
		for n%divisor == 0 {
			factors = append(factors, divisor)
			n /= divisor
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}

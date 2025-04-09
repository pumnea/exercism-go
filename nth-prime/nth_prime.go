// Package prime provides utilities to find nth prime number.
package prime

import "errors"

// Nth returns the nth prime number.
// An error is returned if the nth prime number can't be calculated
// Or if 'n' is equal or less than zero.
func Nth(n int) (int, error) {
	// Early return if n is negative
	if n <= 0 {
		return 0, errors.New("n cannot be equal or less than zero")
	}

	// Early return since 1st prime is 2
	if n == 1 {
		return 2, nil
	}

	// Count for generated prime numbers
	count := 1

	// Next candidate prime number
	candidate := 3

	for {
		if isPrime(candidate) {
			count++
			if count == n {
				return candidate, nil
			}
		}
		candidate += 2
	}
}

// isPrime checks if a number is prime
func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}

	// Check odd divisors up to square root of num
	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

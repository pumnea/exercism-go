// Package collatzconjecture computes properties of the Collatz sequence.
package collatzconjecture

import (
	"errors"
)

// CollatzConjecture returns the steps to reach 1 from n via the Collatz sequence, or an error if n <= 0.
func CollatzConjecture(n int) (int, error) {
	count := 0
	res := n
	if res <= 0 {
		return count, errors.New("input must be positive")
	}
	if res == 1 {
		return count, nil
	}
	for res != 1 {
		if res%2 == 0 {
			res /= 2
		} else {
			res = res*3 + 1
		}
		count++
	}
	return count, nil
}

// Package grains calculates the number of grains on a chessboard using the wheat and chessboard problem.
// In this problem, grains are placed on a chessboard such that the first square has 1 grain,
// the second square has 2 grains, the third square has 4 grains, and so on, doubling each time.
package grains

import (
	"errors"
)

// / Square calculates the number of grains on a specific square of the chessboard.
// The first square has 1 grain, the second 2, the third 4, etc.
// Returns an error if the square number is outside the valid range of 1-64.
func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("number is out of the range 1..64")
	}
	return 1 << (number - 1), nil
}

// Total calculates the total number of grains on the entire chessboard.
// This is the sum of grains on all 64 squares, which is 2^64 - 1.
func Total() uint64 {
	return (1 << 64) - 1
}

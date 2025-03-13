package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fc FodderCalculator, num int) (float64, error) {
	total, error := fc.FodderAmount(num)
	if error != nil {
		return 0, error
	}
	ff, error := fc.FatteningFactor()
	if error != nil {
		return 0, error
	}
	return total * ff / float64(num), error
}

func ValidateInputAndDivideFood(fc FodderCalculator, num int) (float64, error) {
	if num <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, num)
}

type InvalidCowsError struct {
	num     int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d %s", e.num, e.message)
}

func ValidateNumberOfCows(num int) error {
	if num < 0 {
		return &InvalidCowsError{
			num:     num,
			message: "cows are invalid: there are no negative cows",
		}
	}
	if num == 0 {
		return &InvalidCowsError{
			num:     num,
			message: "cows are invalid: no cows don't need food",
		}
	}
	return nil
}

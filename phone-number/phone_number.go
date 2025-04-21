// Package phonenumber provides utilities to clean up user-entered phone numbers.
// It handles NANP (North American Numbering Plan) numbers.
package phonenumber

import (
	"errors"
	"fmt"
)

var (
	n11Codes = map[string]bool{
		"211": true, "311": true, "411": true, "511": true,
		"611": true, "711": true, "811": true, "911": true,
	}
	// errNoDigits is returned when phoneNumber is empty.
	errNoDigits = errors.New("no digits found in input")
	// errInvalidLength is returned when the phone number has an incorrect number of digits.
	errInvalidLength = errors.New("phone number must have 10 or 11 digits")
	// errInvalidCountryCode is returned when an 11-digit number does not start with country code 1.
	errInvalidCountryCode = errors.New("invalid country code")
	// errInvalidAreaCode is returned when the area code is invalid.
	errInvalidAreaCode = errors.New("invalid area code")
	// errInvalidExchangeCode is returned when the exchange code is invalid.
	errInvalidExchangeCode = errors.New("exchange code must start with 2–9 and not be an N11 code")
)

// Number cleans a phone number by removing non-digits and validates it as an NANP number.
// It returns a 10-digit string (area code + exchange + subscriber) or an error.
// The input must be:
// - 10 digits (NPA + NXX + XXXX), or
// - 11 digits starting with country code 1.
// The area code and exchange code must start with 2–9 and not be N11 codes (e.g., 911).
func Number(phoneNumber string) (string, error) {
	digits, err := cleanPhoneNumber(phoneNumber)
	if err != nil {
		return "", err
	}

	if len(digits) != 10 && len(digits) != 11 {
		return "", fmt.Errorf("%w: got %d digits", errInvalidLength, len(digits))
	}

	if len(digits) == 11 {
		if digits[0] != '1' {
			return "", fmt.Errorf("%w: got %c", errInvalidCountryCode, digits[0])
		}
		// remove country code
		digits = digits[1:]
	}

	areaCode := digits[:3]
	if !isValidCode(areaCode) {
		return "", fmt.Errorf("%w: %s", errInvalidAreaCode, areaCode)
	}

	exchangeCode := digits[3:6]
	if !isValidCode(exchangeCode) {
		return "", fmt.Errorf("%w: %s", errInvalidExchangeCode, exchangeCode)
	}

	return digits, nil
}

func AreaCode(phoneNumber string) (string, error) {
	digits, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	// extract the area code (first 3 digits)
	return digits[:3], nil
}

func Format(phoneNumber string) (string, error) {
	digits, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	// Format as (NPA) NXX-XXXX.
	return fmt.Sprintf("(%s) %s-%s", digits[:3], digits[3:6], digits[6:]), nil
}

// cleanPhoneNumber removes non-digit characters from a phone number.
// It returns the cleaned string or an error if no digits are found.
func cleanPhoneNumber(phoneNumber string) (string, error) {
	digits := make([]byte, 0, 11)
	for _, r := range phoneNumber {
		if r >= '0' && r <= '9' {
			digits = append(digits, byte(r))
		}
	}
	if len(digits) == 0 {
		return "", errNoDigits
	}

	return string(digits), nil
}

// isValidCode checks if a 3-digit code is valid for NANP area or exchange codes.
// It must start with 2–9 and not be an N11 code (e.g., 911, 411).
func isValidCode(code string) bool {
	if len(code) != 3 || code[0] < '2' || code[0] > '9' {
		return false
	}
	return !isN11Code(code)
}

// isN11Code checks if the code is an N11 service code (e.g., 211, 911).
func isN11Code(code string) bool {
	return n11Codes[code]
}

// Package cipher provides utilities to encrypt using caesar cipher.
package cipher

import (
	"strings"
	"unicode"
)

type shift int

type vigenere struct {
	key string
}

// NewCaesar creates a new Caesar cipher with a fixed shift of 3.
func NewCaesar() Cipher {
	return shift(3)
}

// NewShift creates a new shift cipher with the given distance.
// Returns nil if distance is 0 or outside the range [-25, 25].
func NewShift(distance int) Cipher {
	if distance == 0 || distance < -25 || distance > 25 {
		return nil
	}
	return shift(distance)
}

// Encode encrypts the input string using a shift cipher.
func (c shift) Encode(input string) string {
	var result strings.Builder
	for _, r := range strings.ToLower(input) {
		if unicode.IsLetter(r) {
			// Apply shift within a-z range (97-122 in ASCII)
			shifted := (int(r-'a') + int(c)) % 26
			if shifted < 0 {
				shifted += 26
			}
			result.WriteByte(byte(shifted + 'a'))
		}
	}
	return result.String()
}

// Decode decrypts the input string using a shift cipher.
func (c shift) Decode(input string) string {
	return shift(-c).Encode(input)
}

// NewVigenere creates a new Vigenere cipher with the given key.
// Returns nil if the key is empty, contains non-letter characters,
// or consists entirely of the letter 'a'.
func NewVigenere(key string) Cipher {
	if key == "" {
		return nil
	}

	allA := true
	for _, r := range key {
		if !unicode.IsLower(r) || r < 'a' || r > 'z' {
			return nil
		}
		if r != 'a' {
			allA = false
		}
	}

	if allA {
		return nil
	}

	return vigenere{key: key}
}

// Encode encrypts the input string using a Vigenere cipher.
func (v vigenere) Encode(input string) string {
	var result strings.Builder
	keyIndex := 0
	keyLen := len(v.key)

	for _, r := range strings.ToLower(input) {
		if unicode.IsLetter(r) {
			// Calculate shift from key character (a=0, b=1, etc.)
			keyShift := int(v.key[keyIndex%keyLen] - 'a')

			// Apply shift within a-z range
			shifted := (int(r-'a') + keyShift) % 26
			result.WriteByte(byte(shifted + 'a'))

			keyIndex++
		}
	}
	return result.String()
}

// Decode decrypts the input string using a Vigenere cipher.
func (v vigenere) Decode(input string) string {
	var result strings.Builder
	keyIndex := 0
	keyLen := len(v.key)

	for _, r := range input {
		if unicode.IsLetter(r) {
			// Calculate reverse shift from key character
			keyShift := int(v.key[keyIndex%keyLen] - 'a')

			// Apply reverse shift within a-z range
			shifted := (int(r-'a') - keyShift) % 26
			if shifted < 0 {
				shifted += 26
			}
			result.WriteByte(byte(shifted + 'a'))

			keyIndex++
		}
	}
	return result.String()
}

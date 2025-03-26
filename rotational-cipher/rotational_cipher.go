// Package rotationalcipher provides utility to encrypt a phrase using rotational cypher.
package rotationalcipher

// RotationalCipher encrypts a string using a Caesar cipher with the given shift key.
// The shiftKey determines how many positions each alphabetic character (a-z, A-Z) is shifted forward.
// Non-alphabetic characters remain unchanged. The shift wraps around the alphabet (e.g., 'z' + 1 = 'a').
// Returns the encrypted result as a string.
func RotationalCipher(plain string, shiftKey int) string {
	// Return plain unchanged if it's empty or shiftKey is zero
	if plain == "" || shiftKey == 0 {
		return plain
	}

	// Normalize shiftKey to 0-25
	shiftKey = shiftKey % 26
	if shiftKey < 0 {
		shiftKey += 26
	}
	// Pre-allocate capacity for efficiency
	result := make([]rune, 0, len(plain))
	for _, r := range plain {
		switch {
		case r >= 'a' && r <= 'z':
			// Shift lowercase letters
			// Compute position in alphabet (0-25); 'a' = 97, 'b' = 98, etc., in ASCII/Unicode
			pos := int(r - 'a')
			// Calculate new position with wraparound
			newPos := (pos + shiftKey) % 26
			result = append(result, 'a'+rune(newPos))
		case r >= 'A' && r <= 'Z':
			// Shift uppercase letters
			// Compute position in alphabet (0-25); 'A' = 65, 'B' = 66, etc., in ASCII/Unicode
			pos := int(r - 'A')
			// Calculate new position with wraparound
			newPos := (pos + shiftKey) % 26
			result = append(result, 'A'+rune(newPos))
		default:
			// Escape non-letters
			result = append(result, r)
		}
	}
	return string(result)
}

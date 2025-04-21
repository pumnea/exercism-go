// Package series provides utilities to extract contiguous substrings of a given length from a string.
// All returns all substrings, while UnsafeFirst returns the first substring.
// For example, All(3, "49142") returns ["491", "914", "142"], and UnsafeFirst(3, "49142") returns "491".
package series

// All returns all contiguous substrings of length n from s in order.
// If n <= 0 or n > len(s), it returns nil.
// Example:
//
//	All(3, "49142") // Returns ["491", "914", "142"]
func All(n int, s string) []string {
	if n <= 0 || n > len(s) {
		return nil
	}

	result := make([]string, 0, len(s)-n+1)
	for i := 0; i+n <= len(s); i++ {
		result = append(result, s[i:i+n])
	}
	return result
}

// UnsafeFirst returns the first contiguous substring of length n from s.
// If n <= 0 or n > len(s), it returns an empty string.
// The "unsafe" name indicates it does not return an error for invalid inputs.
// Example:
//
//	UnsafeFirst(3, "49142") // Returns "491"
func UnsafeFirst(n int, s string) string {
	if n <= 0 || n > len(s) {
		return ""
	}
	return s[:n]
}

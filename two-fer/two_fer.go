// Package twofer generates "two for" sharing phrases.
package twofer

import "fmt"

// If name is present final string contains name, otherwise it returns a generic string.
func ShareWith(name string) string {
	if len(name) == 0 {
		return "One for you, one for me."
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}

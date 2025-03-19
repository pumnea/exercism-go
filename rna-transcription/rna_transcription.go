// Package strand provides utility to determine the RNA complement of a given DNA sequence.
package strand

import "strings"

// ToRNA converts a DNA sequence to its RNA complement.
// It replaces A with U, T with A, C with G, and G with C.
func ToRNA(dna string) string {
	var builder strings.Builder
	builder.Grow(len(dna))
	for _, r := range dna {
		switch r {
		case 'A':
			builder.WriteRune('U')
		case 'C':
			builder.WriteRune('G')
		case 'G':
			builder.WriteRune('C')
		case 'T':
			builder.WriteRune('A')
		}
	}
	return builder.String()
}

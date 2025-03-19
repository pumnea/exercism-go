// package dna validates DNA and counts nucleotides.
package dna

import (
	"errors"
	"strings"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

var validNucleotides = map[rune]bool{
	'A': true, 'C': true, 'G': true, 'T': true,
}

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	if d == "" {
		return h, nil
	}
	for _, r := range strings.ToUpper(string(d)) {
		if !validNucleotides[r] {
			return nil, errors.New("not a valid DNA sequence")
		}

		h[r]++
	}
	return h, nil
}

// Package dna validates DNA and counts nucleotides.
package dna

import (
	"errors"
	"strings"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

// ErrInvalidNucleotide is returned when an invalid nucleotide is encountered.
var ErrInvalidNucleotide = errors.New("invalid nucleotide in DNA sequence")

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	// Initialize histogram with all valid nucleotides at zero.
	h := Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}

	// Handle empty input
	if len(d) == 0 {
		return h, nil
	}

	// Count nucleotides
	for _, r := range strings.ToUpper(string(d)) {
		switch r {
		case 'A', 'C', 'G', 'T':
			h[r]++
		default:
			return nil, ErrInvalidNucleotide
		}
	}
	return h, nil
}

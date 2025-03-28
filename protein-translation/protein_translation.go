// Package protein provides utilities to translate RNA sequences into protein sequences.
package protein

import "errors"

// ErrStop represents a stop codon encountered during translation.
var ErrStop = errors.New("stop codon")

// ErrInvalidBase represents an error for an invalid nucleotide base or codon.
var ErrInvalidBase = errors.New("invalid base")

// codonToProtein maps codons to their corresponding amino acids.
var codonToProtein = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "", // Stop codons return empty string with ErrStop
	"UAG": "",
	"UGA": "",
}

// FromRNA translates an RNA sequence into a protein sequence.
// It breaks the RNA into codons (3 nucleotides each) and translates them into amino acids,
// stopping at the first stop codon (UAA, UAG, UGA). Returns an error if the RNA length
// is not a multiple of 3 or contains an invalid base.
func FromRNA(rna string) ([]string, error) {
	// Check if RNA length is a multiple of 3
	if len(rna)%3 != 0 {
		return nil, ErrInvalidBase
	}

	// Resulting protein sequence
	protein := make([]string, 0, len(rna)/3)

	// Process each codon
	for i := 0; i < len(rna); i += 3 {
		codon := rna[i : i+3]
		amino, err := FromCodon(codon)
		if err == ErrStop {
			return protein, nil // Stop codon terminates normally
		}
		if err != nil {
			return nil, err // Propagate invalid base error
		}
		protein = append(protein, amino)
	}

	return protein, nil
}

// FromCodon translates a single codon into its corresponding amino acid.
// Returns the amino acid name (e.g., "Methionine") or an empty string with ErrStop for stop codons.
// Returns ErrInvalidBase for invalid codons.
func FromCodon(codon string) (string, error) {
	if protein, ok := codonToProtein[codon]; ok {
		if protein == "" {
			return "", ErrStop // Stop codons (UAA, UAG, UGA)
		}
		return protein, nil
	}
	return "", ErrInvalidBase
}

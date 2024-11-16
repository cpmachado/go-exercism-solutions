package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

func newHistogram() Histogram {
	return map[rune]int{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}
}

func validNucleotide(r rune) bool {
	switch r {
	case 'A', 'C', 'G', 'T':
		return true
	default:
		return false
	}
}

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
// /
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	var h Histogram = newHistogram()

	for _, r := range d {
		if validNucleotide(r) {
			h[r]++
		} else {
			return nil, errors.New("Invalid nucleotide")
		}
	}
	return h, nil
}

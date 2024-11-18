package protein

import "errors"

var ErrStop = errors.New("Error Codon")
var ErrInvalidBase = errors.New("Invalid Base")

func FromRNA(rna string) ([]string, error) {
	var ret []string
	for i := 0; i < len(rna); i += 3 {
		v, err := FromCodon(rna[i : i+3])
		switch err {
		case nil:
			ret = append(ret, v)
		case ErrStop:
			return ret, nil
		case ErrInvalidBase:
			return nil, err
		}
	}
	return ret, nil
}

func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	}

	return "", ErrInvalidBase
}

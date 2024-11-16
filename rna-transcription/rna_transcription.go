package strand

func ToRNA(dna string) string {
	ret := make([]rune, len(dna))

	for i, nucleotide := range dna {
		switch nucleotide {
		case 'C':
			ret[i] = 'G'
		case 'G':
			ret[i] = 'C'
		case 'T':
			ret[i] = 'A'
		case 'A':
			ret[i] = 'U'
		}
	}

	return string(ret)
}

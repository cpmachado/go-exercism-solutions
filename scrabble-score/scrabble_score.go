package scrabble

import "strings"

func Score(word string) int {
	sum := 0
	for _, letter := range strings.ToUpper(word) {
		switch letter {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			sum++
		case 'D', 'G':
			sum += 2
		case 'B', 'C', 'M', 'P':
			sum += 3
		case 'F', 'H', 'V', 'W', 'Y':
			sum += 4
		case 'K':
			sum += 5
		case 'J', 'X':
			sum += 8
		case 'Q', 'Z':
			sum += 10
		}
	}
	return sum
}

package luhn

import (
	"strings"
	"unicode"
)

func Valid(id string) bool {
	trimmed := strings.ReplaceAll(id, " ", "")
	n := len(trimmed)
	parity := n % 2
	sum := 0
	if n < 2 {
		return false
	}

	for i, digit := range trimmed {
		// validation
		if !unicode.IsDigit(digit) {
			return false
		}
		v := int(digit - '0')

		// computation
		if i%2 != parity {
			sum += v
		} else if v > 4 {
			sum += 2*v - 9
		} else {
			sum += 2 * v
		}
	}
	return sum%10 == 0
}

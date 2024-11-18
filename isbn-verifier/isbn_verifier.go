package isbn

import (
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	trimmed := strings.ReplaceAll(isbn, "-", "")
	if len(trimmed) != 10 {
		return false
	}
	sum := 0

	for i, ch := range trimmed {
		switch {
		case unicode.IsDigit(ch):
			sum += int(ch-'0') * (10 - i)
		case i == 9 && unicode.ToUpper(ch) == 'X':
			sum += 10
		default:
			return false
		}
	}

	return sum%11 == 0
}

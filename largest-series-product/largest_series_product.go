package lsproduct

import (
	"errors"
	"strings"
	"unicode"
)

func validate(digits string, span int) bool {
	return span < 0 || len(digits) < span || strings.IndexFunc(digits, func(r rune) bool { return !unicode.IsDigit(r) }) >= 0
}

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if validate(digits, span) {
		return 0, errors.New("Invalid input")
	}

	digits_int := make([]int64, len(digits))

	for i, r := range digits {
		digits_int[i] = int64(r - '0')
	}
	largest := int64(0)
	n := len(digits_int) - span + 1

	for i := 0; i < n; i++ {
		prod := digits_int[i]
		for j := 1; j < span; j++ {
			prod *= digits_int[i+j]
		}
		if prod > largest {
			largest = prod
		}
	}

	return largest, nil
}

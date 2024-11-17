package romannumerals

import "errors"

type RomanNumber struct {
	value int
	rep   string
}

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3999 {
		return "", errors.New("Invalid")
	}
	ret := ""
	mappers := []RomanNumber{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	for _, mapper := range mappers {
		for mapper.value <= input {
			ret += mapper.rep
			input -= mapper.value
		}
	}

	return ret, nil
}

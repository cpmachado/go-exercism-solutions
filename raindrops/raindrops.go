package raindrops

import "strconv"

type IntStringTuple struct {
	factor int
	value  string
}

func Convert(number int) string {
	ret := ""
	conversions := []IntStringTuple{
		IntStringTuple{3, "Pling"},
		IntStringTuple{5, "Plang"},
		IntStringTuple{7, "Plong"},
	}
	for _, v := range conversions {
		if number%v.factor == 0 {
			ret += v.value
		}
	}
	if ret == "" {
		ret = strconv.Itoa(number)
	}
	return ret
}

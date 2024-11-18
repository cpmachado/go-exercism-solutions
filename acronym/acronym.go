package acronym

import (
	"regexp"
	"unicode"
)

func Abbreviate(s string) string {
	ret := ""
	re := regexp.MustCompile(`(\s|-|_)`)
	for _, word := range re.Split(s, -1) {
		if word == "" {
			continue
		}
		ret += string(unicode.ToUpper([]rune(word)[0]))
	}
	return ret
}

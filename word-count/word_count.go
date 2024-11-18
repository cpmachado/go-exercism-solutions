package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	ret := Frequency{}
	re := regexp.MustCompile(`([^\w']|\B'|'\B)`)
	phrase = strings.ToLower(phrase)

	for _, word := range re.Split(phrase, -1) {
		if word == "" {
			continue
		}
		strings.ToLower(word)
		ret[word]++
	}
	return ret
}

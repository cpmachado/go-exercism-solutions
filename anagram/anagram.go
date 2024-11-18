package anagram

import "strings"

type RuneCounter map[rune]int

func Detect(subject string, candidates []string) []string {
	var ret []string
	subject = strings.ToUpper(subject)
	subjectLen := len(subject)
	subjectCounter := runeCounter(subject)

	for i, candidate := range candidates {
		candidate = strings.ToUpper(candidate)
		if subjectLen != len(candidate) || subject == candidate {
			continue
		}
		candidateCounter := runeCounter(candidate)
		if counterMatch(subjectCounter, candidateCounter) {
			ret = append(ret, candidates[i])
		}

	}
	return ret
}

func runeCounter(s string) RuneCounter {
	ret := make(RuneCounter)

	for _, r := range s {
		ret[r] += 1
	}
	return ret
}

func counterMatch(a RuneCounter, b RuneCounter) bool {
	for k, v := range a {
		if bv, found := b[k]; !found || v != bv {
			return false
		}
	}
	return true
}

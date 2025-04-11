package wordy

import (
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	if !strings.HasPrefix(question, "What is") {
		return 0, false
	}

	s := strings.TrimPrefix(question, "What is ")
	s = strings.TrimRight(s, "?")

	v := strings.Split(s, " ")

	if len(v) == 0 {
		return 0, false
	}

	acc, err := strconv.Atoi(v[0])
	if err != nil {
		return 0, false
	}

	op := ""
	expectDigit := false
	for _, t := range v[1:] {
		if expectDigit {
			a, err := strconv.Atoi(t)
			if err != nil {
				return 0, false
			}
			switch op {
			case "plus":
				acc += a
			case "minus":
				acc -= a
			case "divided":
				acc /= a
			case "multiplied":
				acc *= a
			default:
				return 0, false
			}
			expectDigit = false
		} else {
			switch t {
			case "plus", "minus":
				expectDigit = true
				fallthrough
			case "divided", "multiplied":
				op = t
			case "by":
				if op != "divided" && op != "multiplied" {
					return 0, false
				}
				expectDigit = true
			default:
				return 0, false
			}
		}
	}

	if expectDigit {
		return 0, false
	}

	return acc, true
}

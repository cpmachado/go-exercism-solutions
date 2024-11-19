package bottlesong

import (
	"errors"
	"fmt"
)

var numberMapper = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
}

func Recite(startBottles, takeDown int) []string {
	if startBottles < 1 {
		return []string{}
	}

	stanzas := [][]string{}
	for i := 0; i < startBottles && i < takeDown; i++ {
		stanza, err := stanza(startBottles - i)
		if err != nil {
			return []string{}
		}
		stanzas = append(stanzas, stanza)
	}

	n := len(stanzas)
	if n == 0 {
		return []string{}
	}

	ret := stanzas[0]
	for i := 1; i < len(stanzas); i++ {
		ret = append(ret, "")
		ret = append(ret, stanzas[i]...)
	}
	return ret
}

func stanza(k int) ([]string, error) {
	switch {
	case k < 1 || k > 10:
		return nil, errors.New("Invalid stanza")
	case k == 1:
		return []string{
			"One green bottle hanging on the wall,",
			"One green bottle hanging on the wall,",
			"And if one green bottle should accidentally fall,",
			"There'll be no green bottles hanging on the wall.",
		}, nil
	case k == 2:
		return []string{
			"Two green bottles hanging on the wall,",
			"Two green bottles hanging on the wall,",
			"And if one green bottle should accidentally fall,",
			"There'll be one green bottle hanging on the wall.",
		}, nil
	default:
		return []string{
			fmt.Sprintf("%s green bottles hanging on the wall,", Title(numberMapper[k-1])),
			fmt.Sprintf("%s green bottles hanging on the wall,", Title(numberMapper[k-1])),
			"And if one green bottle should accidentally fall,",
			fmt.Sprintf("There'll be %s green bottles hanging on the wall.", numberMapper[k-2]),
		}, nil
	}
}

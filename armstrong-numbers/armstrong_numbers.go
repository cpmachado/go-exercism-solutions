package armstrong

import (
	"math"
)

func IsNumber(n int) bool {
	switch {
	case n < 0:
		return false
	case n < 10:
		return true
	case n < 100:
		return false
	}
	p := 0
	for d := n; d > 0; d /= 10 {
		p++
	}

	sum := 0

	for d := n; d > 0; d /= 10 {
		sum += int(math.Pow(float64(d%10), float64(p)))
	}
	return sum == n
}

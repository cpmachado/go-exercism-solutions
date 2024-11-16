package darts

func Score(x, y float64) int {
	squared := x*x + y*y
	switch {
	case squared > 100:
		return 0
	case squared > 25:
		return 1
	case squared > 1:
		return 5
	default:
		return 10
	}
}

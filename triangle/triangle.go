package triangle

type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

func triangleInequality(a, b, c float64) bool {
	return a > 0 && b > 0 && c > 0 && (a+b >= c) && (c+a >= b) && (b+c >= a)
}

func KindFromSides(a, b, c float64) Kind {
	switch {
	case !triangleInequality(a, b, c):
		return NaT
	case a == b && b == c:
		return Equ
	case a == b || a == c || b == c:
		return Iso
	}
	return Sca
}

package complexnumbers

import "math"

// Define the Number type here.
type Number struct {
	Re, Img float64
}

func (n Number) Real() float64 {
	return n.Re
}

func (n Number) Imaginary() float64 {
	return n.Img
}

func (n1 Number) Add(n2 Number) Number {
	return Number{n1.Re + n2.Re, n1.Img + n2.Img}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{n1.Re - n2.Re, n1.Img - n2.Img}
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number{n1.Re*n2.Re - n1.Img*n2.Img, n1.Img*n2.Re + n1.Re*n2.Img}
}

func (n Number) Times(factor float64) Number {
	return Number{factor * n.Re, factor * n.Img}
}

func (n1 Number) Divide(n2 Number) Number {
	return Number{
		(n1.Re*n2.Re + n1.Img*n2.Img) / (math.Pow(n2.Re, 2) + math.Pow(n2.Img, 2)),
		(n1.Img*n2.Re - n1.Re*n2.Img) / (math.Pow(n2.Re, 2) + math.Pow(n2.Img, 2)),
	}
}

func (n Number) Conjugate() Number {
	return Number{n.Re, -n.Img}
}

func (n Number) Abs() float64 {
	return math.Sqrt(math.Pow(n.Re, 2) + math.Pow(n.Img, 2))
}

func (n Number) Exp() Number {
	factor := math.Exp(n.Re)
	return Number{math.Cos(n.Img), math.Sin(n.Img)}.Times(factor)
}

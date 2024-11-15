package diffsquares

func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

func SumOfSquares(n int) int {
	return (2*n*n*n + 3*n*n + n) / 6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

// Package diffsquares provides method for square diffs.
package diffsquares

import "math"

// SquareOfSum squares number sum
func SquareOfSum(i int) int {
	sum := 0
	for sum = 0; i > 0; i-- {
		sum += i
	}
	return int(math.Pow(float64(sum), 2))
}

// SumOfSquares sums number square
func SumOfSquares(i int) int {
	if i == 1 {
		return 1
	}
	return int(math.Pow(float64(i), 2)) + SumOfSquares(i-1)
}

// Difference calculates diff of squares of sum and sum of squares
func Difference(i int) int {
	return SquareOfSum(i) - SumOfSquares(i)
}

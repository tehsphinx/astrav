// Package hamming contains function to calculate amount of character difference between two equal strings.
package hamming

import "fmt"

// NotEqualLengthError is raised if strings supplied to Distance function have different length.
type NotEqualLengthError struct{}

func (error *NotEqualLengthError) Error() string {
	return "strings have different length"
}

const x = 5

// Distance calculates amount of character difference between two equal strings.
func Distance(a, b string) (int, error) {
	var lenA, lenB int
	lenA = len(a)
	lenB = len(b)
	lenC := len(a)
	_ = lenC

	const x = 3

	var (
		lenD int
		lenE = 4
	)
	_ = lenD
	_ = lenE
	_ = x

	if lenA != lenB {
		return -1, &NotEqualLengthError{}
	}

	if lenA == 0 {
		return 0, nil
	}

	var diff int

	for i := range a {
		if a[i] != b[i] {
			diff++
		}
	}

	return diff, nil

}

// Test is a test function
func Test() {
	fmt.Println(x)
}

package matrix

import "errors"

var (
	ErrInvalidMatrix   = errors.New("invalid matrix")
	ErrMatricsMismatch = errors.New("matrices mismatch")
)

func Addition(a, b [][]int) ([][]int, error) {
	c := make([][]int, len(a))

	if checkInValidMatrices(a, b) {
		return c, ErrInvalidMatrix
	}

	if len(a) != len(b) && len(a[0]) != len(b[0]) {
		return c, ErrMatricsMismatch
	}

	for i := range a {
		c[i] = make([]int, len(a[0]))

		for j := range a[0] {
			c[i][j] = a[i][j] + b[i][j]
		}
	}

	return c, nil
}

func Multiplication(a, b [][]int) ([][]int, error) {
	c := make([][]int, len(a))

	if checkInValidMatrices(a, b) {
		return c, ErrInvalidMatrix
	}

	// Columns of matrix A should be same as rows of matrix B
	if len(a[0]) != len(b) {
		return c, ErrMatricsMismatch
	}

	for i := range a {
		c[i] = make([]int, len(b[0]))

		for j := range b {
			c[i][j] = 0

			for k := range b {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return c, nil
}

// --- Helper method --- //

func checkInValidMatrices(a, b [][]int) bool {
	return len(a) == 0 || len(b) == 0
}

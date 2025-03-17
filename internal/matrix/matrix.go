package matrix

import "errors"

var (
	ErrInvalidMatrix   = errors.New("invalid matrix")
	ErrMatricsMismatch = errors.New("matrices mismatch")
)

func Addition(a, b [][]int) ([][]int, error) {
	c := make([][]int, len(a))

	if len(a) == 0 || len(b) == 0 {
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

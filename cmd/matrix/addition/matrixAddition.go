package main

import (
	"code/internal/matrix"
	"fmt"
)

func main() {
	var (
		a [][]int = make([][]int, 0)
		b [][]int = make([][]int, 0)
	)

	a = append(a, []int{1, 2, 3})
	a = append(a, []int{4, 5, 6})
	a = append(a, []int{7, 8, 9})

	b = append(b, []int{11, 12, 13})
	b = append(b, []int{14, 15, 16})
	b = append(b, []int{17, 18, 19})

	fmt.Printf("Matrix A %v\n", a)
	fmt.Printf("Matrix B %v\n", b)

	c, err := matrix.Addition(a, b)
	if err != nil {
		fmt.Printf("Failed to add matrices, error: %s", err.Error())
	} else {
		fmt.Printf("Addition of above matrices %v", c)
	}
}

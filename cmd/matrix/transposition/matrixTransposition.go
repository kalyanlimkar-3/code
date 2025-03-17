package main

import (
	"code/internal/matrix"
	"fmt"
)

func main() {
	var a [][]int = make([][]int, 0)

	a = append(a, []int{1, 2, 3})
	a = append(a, []int{4, 5, 6})
	a = append(a, []int{7, 8, 9})

	fmt.Printf("Matrix A %v\n", a)

	err := matrix.Transposition(a)
	if err != nil {
		fmt.Printf("Failed to transpose above matrix, error: %s", err.Error())
	} else {
		fmt.Printf("Transposition of above matrix %v", a)
	}
}

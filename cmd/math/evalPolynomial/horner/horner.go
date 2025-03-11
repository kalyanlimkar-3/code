package main

import (
	"code/internal/math/polynomial"
	"fmt"
)

func main() {
	x := 2
	a := []int{1, 2, 3, 4}

	fmt.Println("Sum:", polynomial.Horner(x, 0, a))
}

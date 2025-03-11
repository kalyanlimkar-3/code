package main

import (
	"code/internal/fibonacci"
	"fmt"
)

func main() {
	n := 10
	fmt.Printf("Fibonacci of %d: %d", n, fibonacci.Iterative(n))
}

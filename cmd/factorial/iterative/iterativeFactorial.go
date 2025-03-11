package main

import (
	"code/internal/factorial"
	"fmt"
)

func main() {
	n := 0
	fmt.Printf("Factorial of %d: %d", n, factorial.Iterative(n))
}

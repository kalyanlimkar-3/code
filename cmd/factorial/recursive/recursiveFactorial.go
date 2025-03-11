package main

import (
	"code/internal/factorial"
	"fmt"
)

func main() {
	n := 5
	fmt.Printf("Factorial of %d: %d", n, factorial.Recursive(n))
}

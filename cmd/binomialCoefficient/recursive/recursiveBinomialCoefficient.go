package main

import (
	binomialcoefficient "code/internal/binomialCoefficient"
	"fmt"
)

func main() {
	n := 5
	k := 2
	fmt.Printf("Binomial coefficient of %d, %d: %d", n, k, binomialcoefficient.Recursive(n, k))
}

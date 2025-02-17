package main

import (
	"code/internal/permutation"
)

func main() {
	str := "abcd"
	permutation.RecursivePermutation([]byte(str), 0)
}

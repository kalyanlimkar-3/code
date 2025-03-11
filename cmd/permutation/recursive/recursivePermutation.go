package main

import (
	"code/internal/permutation"
)

func main() {
	str := "abcd"
	permutation.Recursive([]byte(str), 0)
}

package main

import (
	"fmt"

	"code/internal/search"
)

func main() {
	var (
		arr        []int = []int{1, 2, 3, 4, 5}
		searchItem int   = 6
	)

	fmt.Println("Input array:", arr, ", Search item:", searchItem)
	found := search.RecursiveBinary(arr, searchItem, 0, len(arr)-1)
	if found == -1 {
		fmt.Println("Not found")
	} else {
		fmt.Println("Found at index:", found)
	}
}

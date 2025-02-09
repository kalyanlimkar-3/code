package main

import (
	"fmt"

	"code/internal/search"
)

func main() {
	var (
		arr        []int = []int{1, 2, 3, 4, 5}
		searchItem int   = 4
	)

	fmt.Println("Input array: ", arr, ", Search item: ", searchItem)
	found := search.Binary(arr, searchItem)
	if found == -1 {
		fmt.Println("Not found")
	} else {
		fmt.Println("Found at index: ", found)
	}
}

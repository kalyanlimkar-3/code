/* Selection Sort

Algorithm:
	for i:= 0 -> n-1
		Examine arr[i] to arr[n-1] and suppose that the smallest integer is at arr[min];
		Interchange arr[i] and arr[min];
*/

package main

import (
	"fmt"

	"code/internal/sort"
	"code/pkg/util"
)

func main() {
	noOfItems := 5
	arr := util.GenerateRandomArray(noOfItems)

	fmt.Println("Input array:", arr)
	sort.Selection(arr)
	fmt.Println("Array after sorting:", arr)
}

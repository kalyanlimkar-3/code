/* Selection Sort

Algorithm:
	for i:= 0 -> n-1
		Examine arr[i] to arr[n-1] and suppose that the smallest integer is at arr[min];
		Interchange arr[i] and arr[min];
*/

package main

import (
	"fmt"

	"code/utility"
)

func main() {
	noOfItems := 5
	arr := utility.GenerateRandomArray(noOfItems)

	fmt.Println("Input array:", arr)

	for i := 0; i < noOfItems-1; i++ {
		min := i

		for j := i + 1; j < noOfItems; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}

		arr[i], arr[min] = arr[min], arr[i]
	}

	fmt.Println("Array after sorting:", arr)
}

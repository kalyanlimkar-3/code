package utility

import "math/rand"

func GenerateRandomArray(size int) []int {
	arr := make([]int, size)

	for i := 0; i < 5; i++ {
		arr[i] = rand.Intn(200) // Returns random number between 0 to 200.
	}

	return arr
}

package search

func Binary(arr []int, search int) int {
	left := 0
	right := len(arr) - 1
	found := -1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == search {
			found = mid
			break
		} else if arr[mid] < search {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return found
}

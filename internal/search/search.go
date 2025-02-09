package search

func Binary(arr []int, searchItem int) int {
	left := 0
	right := len(arr) - 1
	found := -1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == searchItem {
			found = mid
			break
		} else if arr[mid] < searchItem {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return found
}

func RecursiveBinary(arr []int, searchItem, left, right int) int {
	if left <= right {
		mid := (left + right) / 2

		if arr[mid] == searchItem {
			return mid
		} else if arr[mid] < searchItem {
			return RecursiveBinary(arr, searchItem, mid+1, right)
		} else {
			return RecursiveBinary(arr, searchItem, left, mid-1)
		}
	}

	return -1
}

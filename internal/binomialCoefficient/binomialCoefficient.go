package binomialcoefficient

func Recursive(n, k int) int {
	if k > n {
		return 0
	}

	if k == 0 || n == k {
		return 1
	}

	return Recursive(n-1, k-1) + Recursive(n-1, k)
}

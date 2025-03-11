package factorial

func Recursive(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return n * Recursive(n-1)
}

func Iterative(n int) int {
	fact := 1

	for i := 1; i <= n; i++ {
		fact = fact * i
	}

	return fact
}

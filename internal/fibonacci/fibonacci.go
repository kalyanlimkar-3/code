package fibonacci

func Recursive(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return Recursive(n-1) + Recursive(n-2)
}

func Iterative(n int) int {
	n1, n2 := 0, 1

	if n == 0 {
		return n1
	}

	for i := 0; i < n-1; i++ {
		t := n2
		n2 = n1 + n2
		n1 = t
	}

	return n2
}

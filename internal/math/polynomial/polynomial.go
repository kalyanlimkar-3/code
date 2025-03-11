package polynomial

func Horner(x, i int, a []int) int {
	if i == len(a)-1 {
		return a[i]
	}

	return a[i] + x*Horner(x, i+1, a)
}

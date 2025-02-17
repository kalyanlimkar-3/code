package permutation

import "fmt"

func RecursivePermutation(str []byte, idx int) {
	if idx == len(str) {
		fmt.Println(string(str))
		return
	}

	for i := idx; i < len(str); i++ {
		str[i], str[idx] = str[idx], str[i]
		RecursivePermutation(str, idx+1)
		str[i], str[idx] = str[idx], str[i]
	}
}

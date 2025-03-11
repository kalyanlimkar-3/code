package permutation

import "fmt"

func Recursive(str []byte, idx int) {
	if idx == len(str) {
		fmt.Println(string(str))
		return
	}

	for i := idx; i < len(str); i++ {
		str[i], str[idx] = str[idx], str[i]
		Recursive(str, idx+1)
		str[i], str[idx] = str[idx], str[i]
	}
}

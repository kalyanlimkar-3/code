package truthtable

import "fmt"

func TruthTable(b string, n int) {
	if n == 0 {
		fmt.Println(string(b))
		return
	}

	TruthTable(b+"T", n-1)
	TruthTable(b+"F", n-1)
}

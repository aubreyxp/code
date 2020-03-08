package main

import "fmt"

func GetValue(n int, i int) int {
	if n < 1 {
		return 0
	}

	if i < 0 {
		return 0
	}

	if n == 1 && i == 0 {
		return 1
	}

	return GetValue(n-1, i) + GetValue(n-1, i-1)
}

func OutPut(n int) {
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("%d ", GetValue(i, j))
		}
		fmt.Println()
	}
}

func main() {
	OutPut(8)
	fmt.Println("vim-go")
}

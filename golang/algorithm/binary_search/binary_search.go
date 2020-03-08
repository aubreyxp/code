package main

import "fmt"

func find(arrs []int, findNum int) int {
	if len(arrs) == 0 {
		return -1
	}

	firstIndex := 0
	lastIndex := len(arrs) - 1
	for {
		mid := firstIndex + (lastIndex-firstIndex)/2
		if findNum == arrs[mid] {
			return mid
		}

		if lastIndex <= firstIndex {
			return -1
		}

		if findNum > arrs[mid] {
			firstIndex = mid + 1
			continue
		}

		if findNum < arrs[mid] {
			lastIndex = mid - 1
			continue
		}
	}
}

func main() {
	arrs := []int{1, 3, 8, 10}
	index := find(arrs, 1)
	fmt.Println("index:", index)
}

package main

import "fmt"
import "sort"

func QuitSort(arrs []int) {
	if len(arrs) <= 1 {
		return
	}

	midValue := arrs[0]
	left := 0
	right := len(arrs) - 1
	pos := 0

	for {
		if left >= right {
			arrs[pos] = midValue
			break
		}

		for {
			if left >= right {
				break
			}

			if arrs[right] >= midValue {
				right--
				continue
			}

			if arrs[right] < midValue {
				arrs[pos] = arrs[right]
				pos = right
				right--
				break
			}
		}

		for {
			if left >= right {
				break
			}

			if arrs[left] < midValue {
				left++
				continue
			}

			if arrs[left] >= midValue {
				arrs[pos] = arrs[left]
				pos = left
				left++
				break
			}
		}
	}

	if pos > 0 {
		QuitSort(arrs[:pos])
	}

	if pos < len(arrs)-1 {
		QuitSort(arrs[pos+1:])
	}
}

func main() {
	arrs := []int{1, 4, 100, 6, 7, 30, 4, 1}
	QuitSort(arrs)
	sort.Ints(arrs)
	fmt.Println(arrs)
}

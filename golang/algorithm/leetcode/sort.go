package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1
	mid := left + (right-left)/2
	for left < right {
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
		mid = left + (right-left)/2
	}

	if nums[left] == target {
		return left
	} else {
		return -1
	}
}

func main() {
	index := search([]int{-1, 0, 3, 5, 9, 12}, 9)
	fmt.Println(index)
}

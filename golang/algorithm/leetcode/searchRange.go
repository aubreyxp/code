package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	fmt.Println("findFirstTarget")
	left := findFirstTarget(nums, target)
	if left == -1 {
		return []int{-1, -1}
	}

	fmt.Println("findLastTarget")
	right := findLastTarget(nums, target)
	return []int{left, right}
}

func findFirstTarget(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if nums[left] == target {
		return left
	}

	return -1
}

func findLastTarget(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left+1)/2
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid
		}
	}

	return right
}

func main() {
	ret := searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	fmt.Println(ret)
}

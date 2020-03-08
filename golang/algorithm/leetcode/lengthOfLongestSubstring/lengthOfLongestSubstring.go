package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	left := 0
	right := 0
	max := 0
	mp := map[byte]int{}
	for ; right < len(s); right++ {
		cnt, _ := mp[s[right]] // cnt 代表第几个, 从1开始
		if cnt > 0 && cnt > left {
			if right-left > max {
				max = right - left
			}
			left = cnt
		}
		mp[s[right]] = right + 1
	}

	if right-left > max {
		max = right - left
	}

	return max
}

func main() {
	cnt := lengthOfLongestSubstring("au")
	fmt.Println(cnt)
}

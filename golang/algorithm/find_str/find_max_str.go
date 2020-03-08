package main

import "fmt"

func FindFristMaxNoRepeatStr(str string) string {
	if str == "" {
		return str
	}

	maxLen := 0
	maxStr := ""
	info := map[byte]struct{}{}
	firstIndex := 0
	secondIndex := 0
	for i := 0; i < len(str); i++ {
		if _, ok := info[str[i]]; ok {
			if secondIndex-firstIndex+1 > maxLen {
				maxLen = secondIndex - firstIndex + 1
				maxStr = str[firstIndex : secondIndex+1]
				fmt.Println(maxLen, maxStr)
			}
			firstIndex = i
			secondIndex = i
			info = map[byte]struct{}{}
			info[str[i]] = struct{}{}
			continue
		}
		secondIndex = i
		info[str[i]] = struct{}{}
	}

	if maxStr == "" {
		maxStr = str[firstIndex : secondIndex+1]
	}

	return maxStr
}

func main() {
	maxStr := FindFristMaxNoRepeatStr("jiijifefwp")
	fmt.Println(maxStr)
}

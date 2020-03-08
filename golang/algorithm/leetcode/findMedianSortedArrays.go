package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	mids := []int{}
	sum := len1 + len2
	if sum <= 1 {
		mids = append(mids, 1)
	} else {
		if (len1+len2)%2 == 0 {
			mids = append(mids, sum/2, sum/2+1)
		} else {
			mids = append(mids, sum/2+1)
		}
	}

	//fmt.Println(mids)

	count := 0
	ret := []int{}
	for i, j := 0, 0; i < len(nums1) || j < len(nums2); {
		if len(mids) == 0 {
			break
		}

		count++
		hit := false

		if count == mids[0] {
			hit = true
			mids = mids[1:]
		}

		tmp := 0
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] < nums2[j] {
				tmp = nums1[i]
				i++
			} else {
				tmp = nums2[j]
				j++
			}
		} else if i >= len(nums1) {
			tmp = nums2[j]
			j++
		} else if j >= len(nums2) {
			tmp = nums1[i]
			i++
		}

		if hit {
			ret = append(ret, tmp)
		}
	}

	//fmt.Println(ret)

	var all float64
	for _, v := range ret {
		all += float64(v)
	}
	return all / float64(len(ret))
}

func main() {
	mid := findMedianSortedArrays([]int{1, 3}, []int{2})
	fmt.Println(mid)
	fmt.Println("vim-go")
}

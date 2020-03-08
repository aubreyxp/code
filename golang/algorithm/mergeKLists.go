package main

import "fmt"

//* Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	checkOver := func(lists []*ListNode) bool {
		isOver := true
		for _, v := range lists {
			if v != nil {
				isOver = false
			}
		}
		return isOver
	}

	if checkOver(lists) {
		return nil
	}

	var head, last *ListNode
	addValue := func(minIndex int, lists []*ListNode) {
		n := lists[minIndex]
		if head == nil {
			head = n
			last = n
		} else {
			last.Next = n
			last = n
		}
		lists[minIndex] = n.Next
	}

	for !checkOver(lists) {
		minIndex := -1
		//minValue := -1
		isFirst := true
		for i, v := range lists {
			if isFirst && v != nil {
				minIndex = i
				//minValue = v
				isFirst = false
				continue
			}

			if v != nil && v.Val < lists[minIndex].Val {
				minIndex = i
				//minValue = v
				continue
			}
		}

		fmt.Println(minIndex)

		addValue(minIndex, lists)

	}

	last.Next = nil

	return head
}

func makeList(nums [][]int) []*ListNode {
	ret := []*ListNode{}
	for i := 0; i < len(nums); i++ {
		var head, last *ListNode
		for j := 0; j < len(nums[i]); j++ {
			n := &ListNode{
				Val: nums[i][j],
			}

			if head == nil {
				head = n
				last = n
			} else {
				last.Next = n
				last = n
			}
		}

		ret = append(ret, head)
	}

	return ret
}

func dumpList(lists []*ListNode) {
	for _, v := range lists {
		for v != nil {
			fmt.Printf("%d ", v.Val)
			v = v.Next
		}
		fmt.Printf("\n")
	}
}

func dump(lists *ListNode) {
	v := lists
	for v != nil {
		fmt.Printf("%d ", v.Val)
		v = v.Next
	}
	fmt.Printf("\n")

}

func main() {
	//[[1,4,5],[1,3,4],[2,6]]
	ret := makeList([][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}})
	dumpList(ret)
	rr := mergeKLists(ret)
	dump(rr)
}

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stepValue := 0
	var ret *ListNode
	var last *ListNode
	for l1 != nil || l2 != nil {
		l1Value := 0
		l2Value := 0
		if l1 != nil {
			l1Value = l1.Val
		}
		if l2 != nil {
			l2Value = l2.Val
		}
		value := l1Value + l2Value + stepValue
		if value >= 10 {
			stepValue = value / 10
			value = value % 10
		} else {
			stepValue = 0
		}

		ln := &ListNode{
			Val: value,
		}
		if ret == nil {
			ret = ln
			last = ln
		} else {
			last.Next = ln
			last = ln
		}

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	if stepValue != 0 {
		ln := &ListNode{
			Val: stepValue,
		}

		if ret == nil {
			ret = ln
			last = ln
		} else {
			last.Next = ln
			last = ln
		}
	}

	return ret

}

func main() {
	fmt.Println("vim-go")
}

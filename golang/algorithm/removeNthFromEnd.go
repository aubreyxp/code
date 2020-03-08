package main

import "fmt"

//* Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return nil
	}

	first := head
	second := head

	for i := 0; i < n; i++ {
		if second == nil {
			return nil
		}
		second = second.Next
		//fmt.Println("1 step second value :", second.Val)
	}

	var pre *ListNode
	for second != nil {
		second = second.Next
		pre = first
		first = first.Next
	}

	//fmt.Println("pre value :", pre.Val)
	//fmt.Println("first value :", first.Val)

	if pre == nil {
		return head.Next
	} else {
		pre.Next = first.Next
	}

	return head
}

func main() {
	n1 := &ListNode{
		Val: 2,
	}
	n2 := &ListNode{
		Val:  1,
		Next: n1,
	}
	head := removeNthFromEnd(n2, 1)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	fmt.Println("vim-go")
}

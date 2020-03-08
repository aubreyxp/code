package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1 2 3 3 4 4 5
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p1 := head
	p2 := head.Next

	for p2 != nil {
		if (p2.Next == nil) || p2.Next != nil && p2.Val != p2.Next.Val {
			if p2.Val == p1.Val {
				if p1 == head {
					head = p2.Next
					p1 = head
				} else {
					p1.Next = p2.Next
				}
			} else {
				if p1.Next == p2 {
					p1 = p2
				} else {
					p1.Next = p2.Next
				}
			}

			if p1 != nil {
				p2 = p1.Next
			} else {
				p2 = nil
			}
		} else if p2.Next == nil {
		} else {
			p2 = p2.Next
		}
	}

	return head
}

func main() {
	n5 := ListNode{
		Val: 2,
	}
	n6 := ListNode{
		Val:  2,
		Next: &n5,
	}
	n7 := ListNode{
		Val:  1,
		Next: &n6,
	}

	h := deleteDuplicates(&n7)
	cur := h
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}

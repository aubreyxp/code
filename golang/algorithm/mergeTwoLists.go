package main

//* Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var last *ListNode

	if l1 == nil && l2 == nil {
		return nil
	}

	addList := func(n *ListNode) {
		/*
		   new := ListNode{
		           Val: n.Val,
		   }
		*/
		if head == nil {
			head = n
			last = n
		} else {
			last.Next = n
			last = n
		}
	}

	for l1 != nil || l2 != nil {
		if l1 == nil {
			addList(l2)
			l2 = l2.Next
			continue
		}

		if l2 == nil {
			addList(l1)
			l1 = l1.Next
			continue
		}

		if l1.Val < l2.Val {
			addList(l1)
			l1 = l1.Next
		} else {
			addList(l2)
			l2 = l2.Next
		}
	}

	last.Next = nil

	return head
}

func main() {
}

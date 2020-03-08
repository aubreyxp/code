package main

import "fmt"

type Node struct {
	value  int
	height int
	left   *Node
	right  *Node
}

func (n *Node) ChangeHeight() {
	var leftH = -1
	if n.left != nil {
		leftH = n.left.height
	}

	var rightH = -1
	if n.right != nil {
		rightH = n.right.height
	}

	var maxH = leftH
	if rightH > maxH {
		maxH = rightH
	}

	n.height = maxH + 1
}

func (n *Node) Adjust() *Node {
	var leftH = -1
	if n.left != nil {
		leftH = n.left.height
	}

	var rightH = -1
	if n.right != nil {
		rightH = n.right.height
	}

	var diff int
	var childH int
	if rightH > leftH {
		diff = rightH - leftH
		childH = 2
	} else {
		diff = leftH - rightH
		childH = 1
	}

	//fmt.Println("diff:", diff)

	if diff < 2 {
		return nil
	}

	if childH == 2 {
		if insertMark == 1 { // 双
			return n.rmm()
		}

		if insertMark == 2 { // 右
			return n.rr()
		}
	}

	if childH == 1 {
		if insertMark == 1 { // 左
			return n.ll()
		}

		if insertMark == 2 { // 双
			return n.lmm()
		}
	}

	return nil
}

func (n *Node) rmm() *Node {
	//fmt.Println("enter rmm")
	if n.right == nil || n.right.left == nil {
		return nil
	}

	root := n.right.left
	l := root.left
	r := root.right

	root.left = n
	root.right = n.right
	n.right.left = r
	n.right = l

	root.left.ChangeHeight()
	root.right.ChangeHeight()
	root.ChangeHeight()

	return root
}

func (n *Node) lmm() *Node {
	//fmt.Println("enter mm")
	if n.left == nil || n.left.right == nil {
		return nil
	}

	root := n.left.right
	l := root.left
	r := root.right

	root.left = n.left
	root.right = n
	n.left.right = l
	n.left = r

	root.left.ChangeHeight()
	root.right.ChangeHeight()
	root.ChangeHeight()

	return root
}

func (n *Node) ll() *Node {
	//fmt.Println("enter ll")
	if n.left == nil {
		return nil
	}

	right := n.left.right
	left := n.left

	left.right = n
	n.left = right

	n.ChangeHeight()
	left.ChangeHeight()

	return left
}

func (n *Node) rr() *Node {
	//fmt.Println("enter rr")
	if n.right == nil {
		return nil
	}

	left := n.right.left
	right := n.right

	right.left = n
	n.right = left

	n.ChangeHeight()
	right.ChangeHeight()

	return right
}

var insertflag = false
var insertMark = 0

//var lastRoot *Node

func (n *Node) Add(value int) *Node {
	if n == nil {
		n := &Node{
			value:  value,
			height: 0,
		}
		insertflag = true
		return n
	}

	var lastRoot *Node
	if n.value >= value {
		n.left = n.left.Add(value)

		if insertflag {
			insertMark = 1
			insertflag = false
		}

		n.ChangeHeight()
		lastRoot = n.Adjust()
		n.ChangeHeight()
	} else {
		n.right = n.right.Add(value)

		if insertflag {
			insertMark = 2
			insertflag = false
		}

		n.ChangeHeight()
		lastRoot = n.Adjust()
		n.ChangeHeight()
	}

	if lastRoot != nil {
		return lastRoot
	}

	return n
}

func (n *Node) Loop() {
	if n.left != nil {
		n.left.Loop()
	}

	fmt.Println(n.value, n.height)

	if n.right != nil {
		n.right.Loop()
	}
}

func main() {
	var root *Node
	root = root.Add(7)
	root = root.Add(9)
	root = root.Add(8)
	root = root.Add(2)
	root = root.Add(1)
	root = root.Add(0)
	root = root.Add(53)

	root.Loop()
}

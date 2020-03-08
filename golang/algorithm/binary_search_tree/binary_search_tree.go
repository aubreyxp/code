package main

import "fmt"
import "math/rand"

type Elem struct {
	value int
	l     *Elem
	r     *Elem
}

func (e *Elem) print() {
	if e.l != nil {
		e.l.print()
	}
	fmt.Println(e.value)
	if e.r != nil {
		e.r.print()
	}
}

func (e *Elem) Add(value int) {
	if value >= e.value {
		if e.r == nil {
			e.r = &Elem{value: value}
		} else {
			e.r.Add(value)
		}
	}

	if value < e.value {
		if e.l == nil {
			e.l = &Elem{value: value}
		} else {
			e.l.Add(value)
		}
	}
}

type Tree struct {
	root *Elem
}

func (t *Tree) Add(value int) {
	if t.root == nil {
		t.root = &Elem{value: value}
		return
	}

	if value >= t.root.value {
		if t.root.r == nil {
			t.root.r = &Elem{value: value}
		} else {
			t.root.r.Add(value)
		}
	}

	if value < t.root.value {
		if t.root.l == nil {
			t.root.l = &Elem{value: value}
		} else {
			t.root.l.Add(value)
		}
	}
}

func makeBinaryTree() *Tree {
	t := &Tree{}
	for i := 0; i < 10; i++ {
		r := rand.Int()
		t.Add(r)
	}

	return t
}

func loopTree(t *Tree) {
	if t.root.l != nil {
		t.root.l.print()
	}
	fmt.Println(t.root.value)
	if t.root.r != nil {
		t.root.r.print()
	}
}

func main() {
	t := makeBinaryTree()
	loopTree(t)
	fmt.Println("vim-go")
}

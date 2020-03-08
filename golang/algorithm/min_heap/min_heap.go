package main

import "fmt"

type MinHeap struct {
	arrs []int
}

func NewMinHeap(arrs []int) *MinHeap {
	mh := &MinHeap{}
	mh.arrs = make([]int, len(arrs))
	copy(mh.arrs, arrs)

	for i := len(mh.arrs) - 1; i > 0; {
		mh.up(i)
		if i%2 == 0 {
			i -= 2
		} else {
			i--
		}
	}
	return mh
}

func (mh *MinHeap) up(index int) {
	parent := (index - 1) / 2
	left := parent*2 + 1
	right := left + 1
	minIndex := left
	if right < len(mh.arrs) && mh.arrs[right] < mh.arrs[left] {
		minIndex = right
	}
	if mh.arrs[minIndex] < mh.arrs[parent] {
		mh.arrs[minIndex], mh.arrs[parent] = mh.arrs[parent], mh.arrs[minIndex]
		mh.down(minIndex)
	}
}

func (mh *MinHeap) down(index int) {
	left := 2*index + 1
	if left < len(mh.arrs) {
		mh.up(left)
	}
}

func (mh *MinHeap) Add(value int) {
}

func (mh *MinHeap) Pop() (int, bool) {
	if len(mh.arrs) > 0 {
		ret := mh.arrs[0]
		mh.arrs[0] = mh.arrs[len(mh.arrs)-1]
		mh.arrs = mh.arrs[0 : len(mh.arrs)-1]
		mh.down(0)
		return ret, true
	} else {
		return 0, false
	}
}

func main() {
	mp := NewMinHeap([]int{100, 30, 200, 29, 70})
	for {
		value, ok := mp.Pop()
		if !ok {
			break
		}
		fmt.Println(value)
	}
}

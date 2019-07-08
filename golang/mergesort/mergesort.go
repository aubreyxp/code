package main

import (
	"fmt"
	"sort"
	"sync"
)

type Elem struct {
	value int64
	mark  int
}

type MinHeap struct {
	elemList []Elem
	length   int
}

func NewMinHeap(originElemList []Elem) *MinHeap {

	elemLen := len(originElemList)
	mh := &MinHeap{}

	mh.elemList = make([]Elem, elemLen)
	copy(mh.elemList, originElemList)
	mh.length = elemLen

	for i := mh.length; i >= 1; {
		var (
			parent int
			ok     bool
			child  int
		)
		parent = i / 2
		if i%2 == 0 {
			i--
		} else {
			i -= 2
		}
		if parent < 1 {
			continue // root node
		}

		if ok, child = mh.checkAdjust(parent); ok {
			mh.adjustDown(parent, child)
		}
	}

	return mh
}

func (mh *MinHeap) checkMinhp() (ok bool) {
	for i := 1; i <= mh.length; i++ {
		parent := i - 1
		left := i*2 - 1
		right := left + 1
		if left < mh.length && mh.elemList[left].value < mh.elemList[parent].value {
			fmt.Println("left:", left, "parent:", parent)
			fmt.Println("left value:", mh.elemList[left].value)
			fmt.Println("parent value:", mh.elemList[parent].value)
			return false
		}
		if right < mh.length && mh.elemList[right].value < mh.elemList[parent].value {
			fmt.Println("right:", right)
			return false
		}
	}

	return true
}

func (mh *MinHeap) checkAdjust(parent int) (ok bool, child int) {
	if parent > mh.length || parent < 1 {
		return
	}

	left := parent * 2
	right := left + 1

	if left <= mh.length && left >= 1 && mh.elemList[left-1].value < mh.elemList[parent-1].value {
		ok = true
	}

	if right <= mh.length && right >= 1 && mh.elemList[right-1].value < mh.elemList[parent-1].value {
		ok = true
	}

	if !ok {
		return
	}

	child = left
	if right <= mh.length && right >= 1 && mh.elemList[right-1].value < mh.elemList[left-1].value {
		child = right
	}

	return
}

func (mh *MinHeap) adjustPre(parent, child int) {
	var (
		ok           bool
		parentParent int
	)

	for {
		mh.elemList[child-1].value, mh.elemList[parent-1].value = mh.elemList[parent-1].value, mh.elemList[child-1].value
		mh.elemList[child-1].mark, mh.elemList[parent-1].mark = mh.elemList[parent-1].mark, mh.elemList[child-1].mark
		parentParent = parent / 2
		if ok, child = mh.checkAdjust(parentParent); ok {
			parent = parentParent
		} else {
			break
		}
	}
}

func (mh *MinHeap) adjustDown(parent, child int) {
	var (
		ok         bool
		childChild int
	)

	for {
		mh.elemList[child-1].value, mh.elemList[parent-1].value = mh.elemList[parent-1].value, mh.elemList[child-1].value
		mh.elemList[child-1].mark, mh.elemList[parent-1].mark = mh.elemList[parent-1].mark, mh.elemList[child-1].mark
		if ok, childChild = mh.checkAdjust(child); ok {
			parent = child
			child = childChild
		} else {
			break
		}
	}
}

func (mh *MinHeap) Insert(elem Elem) {
	var (
		totalLen   int
		parent     int
		child      int
		ok         bool
		expandList []Elem
	)

	totalLen = len(mh.elemList)
	if mh.length == totalLen {
		expandList = make([]Elem, totalLen*2)
		copy(expandList, mh.elemList)
		mh.elemList = expandList
	}

	mh.elemList[mh.length] = elem
	mh.length++

	parent = (mh.length) / 2
	if ok, child = mh.checkAdjust(parent); ok {
		mh.adjustPre(parent, child)
	}
}

func (mh *MinHeap) Delete() (ok bool, elem Elem) {
	var (
		child  int
		parent int
		isOk   bool
	)

	if mh.length < 1 {
		return
	}

	ok = true
	elem = mh.elemList[0]

	// get last to first, then adjust down
	mh.elemList[0] = mh.elemList[mh.length-1]
	mh.length--

	parent = 1
	if isOk, child = mh.checkAdjust(parent); isOk {
		mh.adjustDown(parent, child)
	}

	return
}

var SplitCount = 1000

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {

	orderList := [][]int64{}
	indexList := []int{}
	heapList := []Elem{}
	lastList := []int64{}
	lock := new(sync.Mutex)

	addOrderList := func(ll []int64) {
		lock.Lock()
		defer lock.Unlock()
		orderList = append(orderList, ll)
	}

	var waitgroup sync.WaitGroup
	totalLen := len(src)
	start := 1
	for ; start <= totalLen; start += SplitCount {

		count := SplitCount
		if start+SplitCount-1 > totalLen {
			count = totalLen - start + 1
		}

		if count <= 0 {
			continue
		}

		end := start - 1 + count

		/*
			fmt.Println("out start:", start)
			fmt.Println("out count:", count)
			fmt.Println("out end:", end)
		*/

		//fmt.Println("out start:", start-1)
		//fmt.Println("out end:", end)

		waitgroup.Add(1)
		go func(start, end int) {

			defer func() {
				waitgroup.Done()
			}()

			//fmt.Println("start:", start)
			//fmt.Println("end:", end)

			expect := src[start:end]
			sort.Slice(expect, func(i, j int) bool { return expect[i] < expect[j] })
			addOrderList(expect)

		}(start-1, end)
	}

	waitgroup.Wait()

	// :最小堆初始化
	for i, l := range orderList {
		elem := Elem{}
		elem.mark = i + 1
		elem.value = l[0]
		indexList = append(indexList, 1)
		heapList = append(heapList, elem)
	}

	mh := NewMinHeap(heapList)
	for {
		var elem Elem
		var ok bool
		if ok, elem = mh.Delete(); !ok {
			break
		}

		lastList = append(lastList, elem.value)
		if len(orderList[elem.mark-1]) > indexList[elem.mark-1] {
			indexList[elem.mark-1]++
			e := Elem{
				mark:  elem.mark,
				value: orderList[elem.mark-1][indexList[elem.mark-1]-1],
			}
			mh.Insert(e)
		}
	}

	copy(src, lastList)
}

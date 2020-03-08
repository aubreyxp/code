package main

import (
	"fmt"
	"strings"
)

// 计算器

type Stack struct {
	record []string
	top    int
}

func NewStack() *Stack {
	return &Stack{
		top:    -1,
		record: []string{},
	}
}

func (s *Stack) Push(r string) {
	if s.top == len(s.record)-1 {
		s.record = append(s.record, r)
		s.top++
		return
	}

	s.top++
	s.record[s.top] = r
	return
}

func (s *Stack) Empty() bool {
	return s.top == -1
}

func (s *Stack) Top() string {
	if s.top == -1 {
		panic("top empty")
	}

	return s.record[s.top]
}

func (s *Stack) Pop() string {
	if s.top == -1 {
		return ""
	}

	ret := s.record[s.top]
	s.top--
	return ret
}

func main() {
	s := NewStack()
	s.Push("1")
	s.Push("1")
	s.Push("1")
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	s.Push("2")
	s.Push("3")
	s.Push("4")
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

func calculate(s string) int {
	if s == "" {
		panic("empty")
	}

	s = strings.Replace(s, " ", "", -1)

	return 0
}

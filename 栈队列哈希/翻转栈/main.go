package main

import (
	"fmt"

	"github.com/zboyco/go-algorithm/gotype"
)

func main() {
	s := &gotype.LinkedStack{}
	for i := 0; i < 10; i++ {
		s.Push(i)
	}
	reverseStack(s)
	for !s.IsEmpty() {
		fmt.Println(s.Pop())
	}
}

func reverseStack(s *gotype.LinkedStack) {
	if s.Size() < 2 {
		return
	}
	top1, _ := s.Pop()
	reverseStack(s)
	top2, _ := s.Pop()
	reverseStack(s)
	s.Push(top1)
	reverseStack(s)
	s.Push(top2)
}

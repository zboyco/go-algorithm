package main

import (
	"fmt"

	"github.com/zboyco/go-algorithm/gotype"
)

func main() {
	s1 := &gotype.LinkedStack{}
	s1.Push(1)
	s1.Push(2)
	s1.Push(3)
	s1.Push(4)
	s1.Push(5)
	fmt.Println(s1.Top())

	s2 := &gotype.LinkedStack{}
	v, _ := s1.Pop()
	for v != nil {
		s2.Push(v)
		v, _ = s1.Pop()
	}
	fmt.Println(s2.Top())
}

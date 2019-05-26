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
	reverseStack1(s)
	for !s.IsEmpty() {
		fmt.Println(s.Pop())
	}
}

func reverseStack1(s *gotype.LinkedStack) {
	// 长度小于2不用翻转
	if s.Size() < 2 {
		return
	}
	// 取出顶元素
	top1, _ := s.Pop()
	// 翻转剩余元素后取出顶元素，也就是翻转前的底元素
	reverseStack1(s)
	top2, _ := s.Pop()

	// 交换取出的两个元素
	reverseStack1(s)
	s.Push(top1)
	reverseStack1(s)
	s.Push(top2)
}

func reverseStack2(s *gotype.LinkedStack) {
	// 长度小于2不用翻转
	if s.Size() < 2 {
		return
	}
	// 取出顶元素
	top, _ := s.Pop()
	// 翻转
	reverseStack2(s)
	// 把取出的元素放入栈底部
	pushToBottom(s, top)
}

func pushToBottom(s *gotype.LinkedStack, v interface{}) {
	// 长度等于0不用处理
	if s.Size() == 0 {
		s.Push(v)
		return
	}
	top, _ := s.Pop()
	pushToBottom(s, v)
	s.Push(top)
}

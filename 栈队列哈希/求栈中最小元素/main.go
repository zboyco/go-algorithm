package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/zboyco/go-algorithm/gotype"
)

func main() {
	s := &gotype.LinkedStack{}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		v := rand.Intn(100) - 50
		fmt.Print(v, " ")
		s.Push(v)
	}
	fmt.Println()
	min, err := calcMin(s)
	if err != nil {
		panic(err)
	}
	fmt.Println("min:", min)
	fmt.Println(s.Top())
}

func calcMin(s *gotype.LinkedStack) (int, error) {
	top, err := s.Top()
	if err != nil {
		return 0, errors.New("empty stack")
	}
	if s.Size() == 1 {
		return top.(int), nil
	}
	s.Pop()                // 弹出栈顶
	min, err := calcMin(s) // 计算栈中剩余元素最小值
	s.Push(top)            // 还原栈顶
	if err != nil {
		return top.(int), nil
	}
	if top.(int) > min {
		return min, nil
	}
	return top.(int), nil
}

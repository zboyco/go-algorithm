package gotype

import (
	"errors"
)

type node struct {
	value interface{}
	next  *node
}

type LinkedStack struct {
	head *node
	size int
}

func (s *LinkedStack) Push(v interface{}) {
	if s.head == nil {
		s.head = &node{}
	}
	newNode := &node{value: v, next: s.head.next}
	s.head.next = newNode
	s.size++
}

func (s *LinkedStack) Size() int {
	return s.size
}

func (s *LinkedStack) Pop() (interface{}, error) {
	if s.head == nil {
		s.head = &node{}
	}
	if s.head.next == nil {
		return nil, errors.New("nil stack")
	}
	result := s.head.next.value
	s.head.next = s.head.next.next
	s.size--
	return result, nil
}

func (s *LinkedStack) Top() (interface{}, error) {
	if s.head == nil {
		s.head = &node{}
	}
	if s.head.next == nil {
		return nil, errors.New("nil stack")
	}
	return s.head.next.value, nil
}

func (s LinkedStack) IsEmpty() bool {
	return s.size == 0
}

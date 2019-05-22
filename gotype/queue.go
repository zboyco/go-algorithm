package gotype

import (
	"errors"
)

type queueNode struct {
	value interface{}
	next  *queueNode
}

type LinkedQueue struct {
	size int
	head *queueNode
	tail *queueNode
}

func (q *LinkedQueue) Enqueue(v interface{}) {
	if q.tail == nil {
		q.tail = &queueNode{value: v}
	} else {
		q.tail.next = &queueNode{value: v}
		q.tail = q.tail.next
	}
	if q.head == nil {
		q.head = q.tail
	}
	q.size++
}

func (q *LinkedQueue) Dequeue() (interface{}, error) {
	if q.head == nil {
		return nil, errors.New("nil queue")
	}
	headNode := q.head
	q.head = q.head.next
	q.size--
	if q.size == 0 {
		q.tail = nil
	}
	return headNode.value, nil
}

func (q *LinkedQueue) GetFront() (interface{}, error) {
	if q.head == nil {
		return nil, errors.New("nil queue")
	}
	return q.head.value, nil
}

func (q *LinkedQueue) GetBack() (interface{}, error) {
	if q.head == nil {
		return nil, errors.New("nil queue")
	}
	return q.tail.value, nil
}

func (q *LinkedQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *LinkedQueue) Size() int {
	return q.size
}

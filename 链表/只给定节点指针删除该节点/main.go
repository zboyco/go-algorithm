package main

import (
	"errors"
	"fmt"
)

type node struct {
	value int
	next  *node
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	list := newLinkedList(s)

	node := list.next.next.next.next.next

	deleteNode2(node)

	printNode(list)
}

func deleteNode(n *node) error {
	if n == nil || n.next == nil {
		return errors.New("nil or last node")
	}
	var pre *node
	for n.next != nil {
		pre = n
		n.value = n.next.value
		n = n.next
	}
	pre.next = nil
	return nil
}

func deleteNode2(n *node) error {
	if n == nil || n.next == nil {
		return errors.New("nil or last node")
	}
	n.value = n.next.value
	n.next = n.next.next
	return nil
}

func newLinkedList(s []int) *node {
	head := &node{}
	cur := head
	for _, v := range s {
		next := &node{
			value: v,
		}
		cur.next = next
		cur = next
	}
	return head
}

func printNode(head *node) {
	next := head.next
	for next != nil {
		fmt.Print(fmt.Sprintf("%v -> ", next.value))
		next = next.next
	}
	fmt.Println()
}

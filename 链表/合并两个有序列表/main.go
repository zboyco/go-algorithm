package main

import (
	"fmt"
)

type node struct {
	value int
	next  *node
}

func main() {
	s1 := []int{0, 1, 3, 6, 7, 9, 12, 13}
	a := newLinkedList(s1)
	s2 := []int{1, 2, 4, 5, 7, 8, 10, 11, 12}
	b := newLinkedList(s2)

	list := mergeList2(a, b)

	printNode(list)
}

// 合并后不影响原链表
func mergeList1(a, b *node) *node {
	head := &node{}
	cur := head
	a = a.next
	b = b.next
	var value int
	for a != nil || b != nil {
		if a != nil && b != nil {
			if a.value < b.value {
				value = a.value
				a = a.next
			} else {
				value = b.value
				b = b.next
			}
		} else if a != nil {
			value = a.value
			a = a.next
		} else {
			value = b.value
			b = b.next
		}
		cur.next = &node{value: value}
		cur = cur.next
	}
	return head
}

// 合并后影响原链表（修改指针）
func mergeList2(a, b *node) *node {
	head := &node{}
	cur := head
	a = a.next
	b = b.next
	for a != nil || b != nil {
		if a != nil && b != nil {
			if a.value < b.value {
				cur.next = a
				a = a.next
			} else {
				cur.next = b
				b = b.next
			}
		} else if a != nil {
			cur.next = a
			a = a.next
		} else {
			cur.next = b
			b = b.next
		}
		cur = cur.next
	}
	return head
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

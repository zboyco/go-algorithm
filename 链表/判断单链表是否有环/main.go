package main

import "fmt"

type node struct {
	value int
	next  *node
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	list := newLinkedList(s)
	last := list.next
	for last.next != nil {
		last = last.next
	}
	last.next = list.next.next.next
	fmt.Println(hasCircle(list))
}

func hasCircle(head *node) bool {
	slow := head.next
	fast := head.next.next
	for slow != nil && fast != nil {
		fmt.Println(slow.value, fast.value)
		if slow == fast {
			return true
		}
		slow = slow.next
		fast = fast.next
		if fast == nil {
			return false
		}
		fast = fast.next
	}
	return false
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

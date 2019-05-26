package main

import "fmt"

type node struct {
	value int
	next  *node
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	list := newLinkedList(s)
	printNode(list)
	n := findnode(list, 3)
	fmt.Println(n.value)
}

func findnode(head *node, k int) *node {
	fast := head
	slow := head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
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

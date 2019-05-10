package main

import "fmt"

type node struct {
	value int
	next  *node
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	list := newLinkedList(s)
	printNode(list)
	reverse(list)
	printNode(list)
}

func reverse(head *node) {
	if head.next == nil {
		return
	}

	cur := head.next.next
	head.next.next = nil
	for cur != nil {
		next := cur.next
		cur.next = head.next
		head.next = cur
		cur = next
		printNode(head) // 打印出翻转步骤
	}
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

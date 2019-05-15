package main

import "fmt"

type node struct {
	value int
	next  *node
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	list := newLinkedList(s)
	printNode(list)
	reverse1(list)
	printNode(list)
}

// ! 逆序
func reverse1(head *node) {
	pre := head
	cur := pre.next
	next := cur.next
	for cur != nil && next != nil {
		cur.next = next.next
		pre.next = next
		next.next = cur
		pre = cur
		cur = pre.next
		next = nil
		if cur != nil {
			next = cur.next
		}
	}
}

// ! 交换数据
func reverse2(head *node) {
	cur := head.next
	next := cur.next
	for cur != nil && next != nil {
		cur.value, next.value = next.value, cur.value
		cur = next.next
		next = nil
		if cur != nil {
			next = cur.next
		}
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

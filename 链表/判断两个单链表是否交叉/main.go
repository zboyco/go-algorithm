package main

import (
	"fmt"
)

type node struct {
	value int
	next  *node
}

func main() {
	s1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	list1 := newLinkedList(s1)
	s2 := []int{9, 10, 11, 12}
	list2 := newLinkedList(s2)

	list2.next.next.next.next.next = list1.next.next.next.next
	printNode(list1)
	printNode(list2)

	fmt.Println(cross3(list1, list2))
}

// * hash法
func cross1(a, b *node) bool {
	tmp := make(map[*node]bool)
	a = a.next
	for a != nil {
		tmp[a] = true
		a = a.next
	}
	b = b.next
	for b != nil {
		if _, ok := tmp[b]; ok {
			return true
		}
		b = b.next
	}
	return false
}

// * 首尾相接法
// * 首先把a链表的尾部链接上b链表的头部
// * 然后判断链接后的链表是否有环即可
func cross2(a, b *node) bool {
	ahead := a
	for a.next != nil {
		a = a.next
	}
	a.next = b.next

	slow := ahead
	fast := ahead
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}
	return false
}

// * 尾节点法
// * 先各自遍历一次记录链表长度
// * 长链表先跳过相差长度，然后一起遍历，如果有相同节点则相交
func cross3(a, b *node) bool {
	aHead, bHead := a, b
	var lenA, lenB int
	for a.next != nil {
		lenA++
		a = a.next
	}
	for b.next != nil {
		lenB++
		b = b.next
	}
	a, b = aHead.next, bHead.next
	if lenA > lenB {
		for lenA != lenB {
			a = a.next
			lenA--
		}
	} else {
		for lenA != lenB {
			b = b.next
			lenB--
		}
	}
	for a != nil {
		if a == b {
			return true
		}
		a = a.next
		b = b.next
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

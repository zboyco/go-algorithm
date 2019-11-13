package main

import "fmt"

func main() {
	fmt.Println("24")
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 非递归实现
func swapPairs(head *ListNode) *ListNode {
	forward := &ListNode{Next: head}
	prev := forward
	curr := forward.Next
	for curr != nil && curr.Next != nil {
		next := curr.Next.Next
		curr.Next.Next = curr
		prev.Next = curr.Next
		curr.Next = next
		prev = curr
		curr = next
	}
	return forward.Next
}

// 递归实现
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	next := head.Next
	head.Next = swapPairs2(next.Next)
	next.Next = head
	return next
}

package main

import "fmt"

func main() {
	fmt.Println("25")
	head := &ListNode{}
	curr := head
	for i := 0; i < 10; i++ {
		curr.Next = &ListNode{
			Val: i + 1,
		}
		curr = curr.Next
	}

	curr = head.Next
	for curr != nil {
		fmt.Print(curr.Val, "->")
		curr = curr.Next
	}

	fmt.Print("\r\n")

	curr = reverseKGroup(head.Next, 3)
	for curr != nil {
		fmt.Print(curr.Val, "->")
		curr = curr.Next
	}
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	curr := head
	for i := 0; i < k; i++ {
		if curr == nil {
			return head
		}
		curr = curr.Next
	}

	prev := &ListNode{Next: head}
	curr = head
	var next *ListNode

	for i := 0; i < k; i++ {
		// if curr == nil {
			
		// }
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	head.Next = reverseKGroup(next, k)

	return prev
}

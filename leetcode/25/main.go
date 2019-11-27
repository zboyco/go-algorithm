package main

import "fmt"

func main() {
	fmt.Println("25")
	head := &ListNode{}
	curr := head
	for i := 0; i < 11; i++ {
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

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var prev *ListNode
	curr := head
	var next *ListNode

	for i := 0; i < k; i++ {
		if curr == nil {
			curr = prev
			prev = nil
			for curr != nil {
				next = curr.Next
				curr.Next = prev
				prev = curr
				curr = next
			}
			return head
		}
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	head.Next = reverseKGroup(next, k)

	return prev
}

// func reverseKGroup(head *ListNode, k int) *ListNode {
// 	curr := head
// 	for i := 0; i < k; i++ {
// 		if curr == nil {
// 			return head
// 		}
// 		curr = curr.Next
// 	}

// 	var prev *ListNode
// 	curr = head
// 	var next *ListNode

// 	for i := 0; i < k; i++ {
// 		next = curr.Next
// 		curr.Next = prev
// 		prev = curr
// 		curr = next
// 	}

// 	head.Next = reverseKGroup(next, k)

// 	return prev
// }

// func reverseKGroup(head *ListNode, k int) *ListNode {
// 	curr := &ListNode{Next: head}
// 	for i := 0; i < k; i++ {
// 		curr = curr.Next
// 		if curr == nil {
// 			return head
// 		}
// 	}

// 	next := curr.Next
// 	curr.Next = nil

// 	newHead := reverseList(head)

// 	head.Next = reverseKGroup(next, k)

// 	return newHead
// }

// func reverseList(head *ListNode) *ListNode {
// 	var prev *ListNode
// 	curr := head
// 	for curr != nil {
// 		next := curr.Next
// 		curr.Next = prev
// 		prev = curr
// 		curr = next
// 	}
// 	return prev
// }

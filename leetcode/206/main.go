package main

import "fmt"

func main() {
	fmt.Println("leetcode 206")
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

	curr = reverseList2(head.Next)
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

// 非递归
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// 递归
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

package main

import "fmt"

type node struct {
	value int
	next  *node
}

func main() {
	s := []int{1, 2, 3, 4, 5, 3, 6, 2, 7, 8, 6, 1, 9, 1}
	list := newLinkedList(s)
	printNode(list)
	// deleteRepeat1(list)
	deleteRepeat2(list)
	printNode(list)
}

// 嵌套循环删除
func deleteRepeat1(head *node) {
	if head == nil {
		return
	}
	next1 := head.next
	for next1 != nil {
		pre2 := next1
		next2 := pre2.next
		for next2 != nil {
			if next1.value == next2.value {
				pre2.next = next2.next
			} else {
				pre2 = next2
			}
			next2 = pre2.next
		}
		next1 = next1.next
	}

}

// 利用map删除
func deleteRepeat2(head *node) {
	if head == nil {
		return
	}

	tmpMap := make(map[int]bool)

	pre := head
	next := head.next
	for next != nil {
		if _, ok := tmpMap[next.value]; ok {
			pre.next = next.next
		} else {
			tmpMap[next.value] = true
			pre = next
		}
		next = pre.next
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

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head := l1
	var tmp int
	for l1 != nil && l2 != nil {
		tmp = l1.Val + l2.Val + tmp
		l1.Val = tmp % 10
		tmp = tmp / 10
		if l1.Next == nil && l2.Next == nil && tmp == 0 {
			break
		}
		if l1.Next == nil {
			l1.Next = &ListNode{
				Val: 0,
			}
		}
		l1 = l1.Next
		if l2.Next == nil {
			l2.Next = &ListNode{
				Val: 0,
			}
		}
		l2 = l2.Next

	}
	return head
}

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	l3 := addTwoNumbers(l1, l2)
	fmt.Println()
	for ; l3 != nil; l3 = l3.Next {
		fmt.Print(l3.Val)
	}
	fmt.Println()
	for ; l1 != nil; l1 = l1.Next {
		fmt.Print(l1.Val)
	}
}

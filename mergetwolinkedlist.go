package main

import "fmt"

type ListNode struct {
	data int
	next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	current := &ListNode{}
	dummy := current
	for l1 != nil && l2 != nil {
		if l1.data <= l2.data {
			dummy.next = l1
			l1 = l1.next
			dummy = dummy.next
		} else {
			dummy.next = l2
			l2 = l2.next
			dummy = dummy.next
		}
	}
	if l1 != nil {
		dummy.next = l1
	}
	if l2 != nil {
		dummy.next = l2
	}
	return current.next
}

func Display(l *ListNode) {
	for l != nil {
		fmt.Printf("%+v ->", l.data)
		l = l.next
	}
	fmt.Println()
}

func main() {

	l1 := &ListNode{
		data: 1, next: &ListNode{
			data: 2,
			next: &ListNode{
				data: 4,
				next: nil,
			},
		},
	}

	l2 := &ListNode{
		data: 5, next: &ListNode{
			data: 6,
			next: &ListNode{
				data: 7,
				next: nil,
			},
		},
	}
	a := mergeTwoLists(l1, l2)
	Display(a)
}

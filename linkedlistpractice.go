package main

type node struct {
	data int
	next *node
}

type linkedList struct {
	len  int
	head *node
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func main() {
	list := linkedList{}
	list.prepend(&node{data: 6})
}

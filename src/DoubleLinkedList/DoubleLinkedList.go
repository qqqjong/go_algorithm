package main

type Node struct {
	data interface{}
	prev *Node
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	len int
}

func (l *LinkedList) InsertAtStart(val interface{}) {
	n := Node{}
	n.data = val
	if l.len == 0 {
		l.head = &n
		l.tail = &n
		l.len++
		return
	}
	n.next = l.head
	n.prev = nil
	l.head.prev = &n
	l.head = &n
	l.len++
}
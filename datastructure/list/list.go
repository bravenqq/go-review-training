// Package list provides ...
package list

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("node not found")
	ErrEmpty    = errors.New("link list is empty")
)

//Node node of list
type Node struct {
	Data int
	next *Node
	pre  *Node
}

//LinkList is a kind of data structure which feature is FIFO
type LinkList struct {
	head *Node
	tail *Node
}

//NewLinkList return a null LinkList
func NewLinkList() *LinkList {
	return &LinkList{head: nil, tail: nil}
}

//Add add a node
func (l *LinkList) Insert(n Node) {
	//when node is the first node of linklist to add
	if l.head == nil {
		l.head = &n
		l.tail = &n
		return
	}
	n.pre = l.tail
	l.tail.next = &n
	l.tail = &n

}

//Remove remove a node
func (l *LinkList) Remove(n Node) error {
	//when n.Data is equal to head.data
	if l.head.Data == n.Data {
		l.head = l.head.next
		return nil
	}
	node := l.head
	for node != nil {
		if node.next.Data == n.Data {
			node.next = node.next.next
			return nil
		}
		node = node.next
	}
	return ErrNotFound
}

//Pop pop a node from the tail
func (l *LinkList) Pop() (Node, error) {
	//when pop all data of linklist
	if l.tail.pre == nil {
		return Node{}, ErrEmpty
	}
	node := l.tail
	l.tail = l.tail.pre
	return *node, nil
}

//Display display all data of LinkList
func (l *LinkList) Display() {
	node := l.head
	for node != nil {
		fmt.Printf("%d\t", node.Data)
		node = node.next
	}
}

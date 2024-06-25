package linkedlist

import (
	"errors"
	"fmt"
)

// Linked List struct. Contains pointer to Head, and Items, containing # of elements
type List struct {
	Head *Node
	Tail *Node
	Items int
}

// a linked list node. Contains the element, Data, and a pointer to the next node, Next
type Node struct {
	Data int
	Next *Node
}

// Append an element to the end of the list
func (l *List) Append(d int) {
	node := &Node{Data: d}
	current := l.Head
	// Check if the list is empty
	if l.Head == nil {
		l.Head = node
	} else {
		l.Tail.Next = node
		current.Next = node
	}
	l.Items++
	l.Tail = node
}

// Add an element at the beginning of the list
func (l *List) Prepend(d int) {
	node := &Node{Data: d}
	node.Next = l.Head
	l.Head = node

	l.Items++
}

// Insert an element into the list after a node.
// Takes int to insert, and *Node to insert after
func (l *List) Insert(d int, n *Node) {
	next := n.Next
	node := &Node{Data: d}
	n.Next =  node
	node.Next = next
	l.Items++
}

// Takes *Node as argument, Removes the node after it
func (l *List) Delete(n *Node) {
	n.Next = n.Next.Next
	l.Items--
}

// Delete first node containing provided value
func (l *List) DeleteValue(d int) error {
	var prev *Node
	current := l.Head
	for current.Next != nil {
		if current.Data == d {
			break
		}
		prev = current
		current = current.Next
	}
	if current.Data == d {
		prev.Next = current.Next
		return nil
	} else {
		err := fmt.Sprintf("item not found: %d\n", d)
		return errors.New(err)
	}

}

func (l *List) Length() int {
	return l.Items
}

// Reverses the linked list
func (l *List) Reverse() {
	var prev, next *Node
	current := l.Head
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
}

// Returns a pointer to the first node containing provided value
func (l *List) Search(d int) *Node {
	current := l.Head
	for current.Data != d {
		current = current.Next
	}
	if current.Data == d {
		return current
	} else {
		return nil
	}
}

func (l *List) isEmpty() bool {
	if l.Head == nil {
		return true
	} else {
		return false
	}
}

func (l *List) Print() {
	current := l.Head
	fmt.Printf("Items: %d | ", l.Items)
	for current.Next != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Printf("%d\n", current.Data)
}
package linkedlist

import (
	"errors"
	"fmt"
)

// List struct. Contains pointer to Head, Tail, and Items, containing # of elements
type List struct {
	Head  *Node
	Tail  *Node
	Items int
}

// Node struct. Contains the element, Data, and a pointer to the next node, Next
type Node struct {
	Data int
	Next *Node
}

// Append an element to the end of the list
func (l *List) Append(d int) {
	node := &Node{Data: d}
	// Check if the list is empty
	if l.Head == nil {
		l.Head = node
	} else {
		l.Tail.Next = node
	}
	l.Tail = node
	l.Items++
}

// Add an element at the beginning of the list
func (l *List) Prepend(d int) {
	node := &Node{Data: d}
	node.Next = l.Head
	l.Head = node
	if l.Tail == nil {
		l.Tail = node
	}
	l.Items++
}

// Insert an element into the list after a node.
// Takes int to insert, and *Node to insert after
func (l *List) Insert(d int, n *Node) {
	if n == nil {
		return
	}
	node := &Node{Data: d}
	node.Next = n.Next
	n.Next = node
	if node.Next == nil {
		l.Tail = node
	}
	l.Items++
}

// Takes *Node as argument, Removes the node after it
func (l *List) Delete(n *Node) error {
	var err error
	if n == nil || n.Next == nil {
		err = fmt.Errorf("node not found: %v", n)
		return err
	}
	n.Next = n.Next.Next
	if n.Next == nil {
		l.Tail = n
	}
	l.Items--
	return err
}

// Delete first node containing provided value.
// Returns nil if item is deleted, else error message
func (l *List) DeleteValue(d int) error {
	if l.Head == nil {
		return errors.New("list is empty")
	}
	if l.Head.Data == d {
		l.Head = l.Head.Next
		if l.Head == nil {
			l.Tail = nil
		}
		l.Items--
		return nil
	}
	var prev *Node
	current := l.Head
	for current != nil && current.Data != d {
		prev = current
		current = current.Next
	}
	if current == nil {
		return fmt.Errorf("item not found: %d", d)
	}
	prev.Next = current.Next
	if prev.Next == nil {
		l.Tail = prev
	}
	l.Items--
	return nil
}

// Length returns the number of elements in the list
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
	if l.Head == nil {
		l.Tail = nil
	} else {
		current = l.Head
		for current.Next != nil {
			current = current.Next
		}
		l.Tail = current
	}
}

// Search returns a pointer to the first node containing provided value
// returns *Node and error - error is nil if item found
func (l *List) Search(d int) (*Node, error) {
	current := l.Head
	var err error

	for current != nil && current.Data != d {
		current = current.Next
	}
	if current == nil {
		err = fmt.Errorf("item not found: %d", d)
	}

	return current, err
}

// isEmpty returns true if the list is empty
func (l *List) isEmpty() bool {
	return l.Head == nil
}

// Print prints the elements of the list
func (l *List) Print() {
	current := l.Head
	fmt.Printf("Items: %d | ", l.Items)
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

func (l *List) Cleanup() {
	l.Head = nil
	l.Tail = nil
	l.Items = 0
}

package linkedlist

import "github.com/jay13jay/metaPractice/mergesort"

// Merge sort the linked list by unpacking into a slice, sorting
// and recreating the linked list
func (l *List) Sort() {
	if l.Head == nil || l.Head.Next == nil {
		return
	}
	// unpack the linked list to a slice
	var listArr []int = make([]int, l.Items)
	current := l.Head
	for i := 0; i < l.Items; i++ {
		listArr[i] = current.Data
		current = current.Next
	}

	// sort the values
	mergesort.MS(listArr)

	// rebuild the list with sorted values
	current = l.Head
	for _, value := range listArr {
		current.Data = value
		current = current.Next
	}
}

// Sort the linked list in place without unpacking into slice

func (l *List) InplaceSort() {
	
}


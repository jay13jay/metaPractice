package linkedlist

import "github.com/jay13jay/metaPractice/mergesort"

// Merge sort the linked list
func (l *List) AtoSort() {
	var listArr []int = make([]int, l.Items)
	current := l.Head
	for i := 0; i < l.Items; i++ {
		listArr[i] = current.Data
		current = current.Next
	}

	mergesort.MS(listArr)
	current = l.Head
	for i := 0; i < l.Items; i++ {
		current.Data = listArr[i]
		current = current.Next
	}

}

func (l *List) InplaceSort() {

}


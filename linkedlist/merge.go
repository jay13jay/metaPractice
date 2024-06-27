package linkedlist

// takes another linked list as an argument
// merges the provided linked list
func (list1 *List) MergeList(list2 *List) *List {
	newList := &List{}
	
	if list1.Head == nil {
		return list2
	}
	if list2.Head == nil {
		return list1
	}

	l1 := list1.Head
	l2 := list2.Head

	// Determine the new head of the merged list
	if l1.Data < l2.Data {
		newList.Head = l1
		l1 = l1.Next
	} else {
		newList.Head = l2
		l2 = l2.Next
	}
	current := newList.Head

	// Merge the lists
	for l1 != nil && l2 != nil {
		if l1.Data < l2.Data {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	// Attach the remaining elements, if any
	if l1 != nil {
		current.Next = l1
		newList.Tail = list1.Tail
	} else {
		current.Next = l2
		newList.Tail = list2.Tail
	}

	// Update tail if new elements were appended
	if newList.Tail == nil {
		for current.Next != nil {
			current = current.Next
		}
		newList.Tail = current
	}
	newList.Items = list1.Items + list2.Items

	return newList
}

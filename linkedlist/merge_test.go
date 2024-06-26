package linkedlist

import (
	"math/rand"
	"testing"
)

// Helper function to create a sorted list of given size with random values
func createSortedList(size int) *List {
	list := &List{}
	for i := 0; i < size; i++ {
		list.Append(rand.Intn(size * 10))
	}
	list.Sort()
	return list
}

// BenchmarkMergeList benchmarks the MergeList method
func BenchmarkMergeList(b *testing.B) {
	for n := 0; n < b.N; n++ {
		l1 := createSortedList(1000)
		l2 := createSortedList(1000)
		l1.MergeList(l2)
	}
}

// TestMergeListCorrectness tests the correctness of the MergeList method
func TestMergeListCorrectness(t *testing.T) {
	tests := []struct {
		list1 []int
		list2 []int
		want  []int
	}{
		{[]int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 1, 1}, []int{1, 1, 1}, []int{1, 1, 1, 1, 1, 1}},
		{[]int{}, []int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{[]int{}, []int{}, []int{}},
	}

	for _, tt := range tests {
		l1 := &List{}
		for _, v := range tt.list1 {
			l1.Append(v)
		}
		l2 := &List{}
		for _, v := range tt.list2 {
			l2.Append(v)
		}
		got := l1.MergeList(l2)

		var gotValues []int
		for node := got.Head; node != nil; node = node.Next {
			gotValues = append(gotValues, node.Data)
		}

		if !equal(gotValues, tt.want) {
			t.Errorf("MergeList() = %v, want %v", gotValues, tt.want)
		}
	}
}

// Helper function to check if two slices are equal
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

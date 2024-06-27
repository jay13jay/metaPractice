package linkedlist

import (
	"math/rand"
	"testing"
)

// BenchmarkAppend benchmarks the Append method
func BenchmarkAppend(b *testing.B) {
	for n := 0; n < b.N; n++ {
		l := &List{}
		for i := 0; i < 1000; i++ {
			l.Append(rand.Intn(1000))
		}
	}
}

// BenchmarkPrepend benchmarks the Prepend method
func BenchmarkPrepend(b *testing.B) {
	for n := 0; n < b.N; n++ {
		l := &List{}
		for i := 0; i < 1000; i++ {
			l.Prepend(rand.Intn(1000))
		}
	}
}

// BenchmarkInsert benchmarks the Insert method
func BenchmarkInsert(b *testing.B) {
	for n := 0; n < b.N; n++ {
		l := &List{}
		l.Append(rand.Intn(1000))
		for i := 0; i < 1000; i++ {
			l.Insert(rand.Intn(1000), l.Head)
		}
	}
}

// BenchmarkDelete benchmarks the Delete method
func BenchmarkDelete(b *testing.B) {
	for n := 0; n < b.N; n++ {
		l := &List{}
		for i := 0; i < 1000; i++ {
			l.Append(rand.Intn(1000))
		}
		for i := 0; i < 1000; i++ {
			l.Delete(l.Head)
		}
	}
}

// BenchmarkDeleteValue benchmarks the DeleteValue method
func BenchmarkDeleteValue(b *testing.B) {
	for n := 0; n < b.N; n++ {
		l := &List{}
		for i := 0; i < 1000; i++ {
			l.Append(rand.Intn(1000))
		}
		for i := 0; i < 1000; i++ {
			l.DeleteValue(rand.Intn(1000))
		}
	}
}

// BenchmarkReverse benchmarks the Reverse method
func BenchmarkReverse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		l := &List{}
		for i := 0; i < 1000; i++ {
			l.Append(rand.Intn(1000))
		}
		l.Reverse()
	}
}

// BenchmarkSearch benchmarks the Search method
func BenchmarkSearch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		l := &List{}
		for i := 0; i < 1000; i++ {
			l.Append(rand.Intn(1000))
		}
		for i := 0; i < 1000; i++ {
			l.Search(rand.Intn(1000))
		}
	}
}

// BenchmarkMerge benchmarks the MergeList method
func BenchmarkMerge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		l1 := &List{}
		l2 := &List{}
		for i := 0; i < 1000; i++ {
			l1.Append(rand.Intn(1000))
			l2.Append(rand.Intn(1000))
		}
		l1.MergeList(l2)
	}
}

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jay13jay/metaPractice/linkedlist"
)

const listSize = 1000000
func main() {
	l1 := &linkedlist.List{}
	l2 := &linkedlist.List{}
	l3 := &linkedlist.List{}

	for i := 0; i < 1000000; i++ {
		insert := rand.Intn(listSize)
		l1.Append(insert)
		insert = rand.Intn(listSize)
		l2.Append(insert)
	}
	// fmt.Println("Pre-sorted lists:")
	// l1.Print()
	// l2.Print()
	fmt.Println("Sorted lists:")
	t1 := time.Now()
	l1.Sort()
	t1a := time.Now()

	t2 := time.Now()
	l2.Sort()
	t2a := time.Now()
	// l1.Print()
	// l2.Print()
	t3 := time.Now()
	l3 = l1.MergeList(l2)
	t3a := time.Now()
	fmt.Printf("l1: %d \tl2: %d \t merged: %d\n", l1.Items, l2.Items, l1.Items + l2.Items)

	fmt.Println("Sort time:")
	fmt.Printf("l1: %s \t", t1a.Sub(t1))
	fmt.Printf("l2: %s\n", t2a.Sub(t2))

	fmt.Printf("Time to merge l1 and l2 (%d items): %v\n", l3.Items, t3a.Sub(t3))
	// fmt.Println("Merged list:")
	l3.Print()

}
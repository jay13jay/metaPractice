package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/jay13jay/metaPractice/linkedlist"
)

const listSize = 10000000
func main() {
	// set up print formatting
	l1 := &linkedlist.List{}
	l2 := &linkedlist.List{}
	l3 := &linkedlist.List{}

	for i := 0; i < listSize; i++ {
		insert := rand.Intn(listSize * 100)
		l1.Append(insert)
		insert = rand.Intn(listSize * 100)
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
	fmt.Printf("l1: %v \tl2: %v \t merged: %v\n", 
		humanize.Comma(int64(l1.Items)), 
		humanize.Comma(int64(l2.Items)), 
		humanize.Comma(int64(l1.Items + l2.Items)))

	fmt.Println("Sort time:")
	fmt.Printf("l1: %s \t", t1a.Sub(t1))
	fmt.Printf("l2: %s\n", t2a.Sub(t2))

	fmt.Printf("Time to merge l1 and l2 (%v items): %v\n", humanize.Comma(int64(l3.Items)), t3a.Sub(t3))
	// fmt.Println("Merged list:")
	// l3.Print()

}
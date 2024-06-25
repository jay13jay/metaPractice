package main

import (
	"fmt"
	"math/rand"

	"github.com/jay13jay/metaPractice/linkedlist"
)

const listSize = 1000
func main() {
	l1 := &linkedlist.List{}
	l2 := &linkedlist.List{}
	l3 := &linkedlist.List{}

	for i := 0; i < 10; i++ {
		insert := rand.Intn(listSize)
		l1.Append(insert)
		insert = rand.Intn(listSize)
		l2.Append(insert)
	}
	fmt.Println("Pre-sorted lists:")
	l1.Print()
	l2.Print()
	fmt.Println("Sorted lists:")
	l1.Sort()
	l2.Sort()
	l1.Print()
	l2.Print()
	l3 = l1.MergeList(l2)
	fmt.Println("Merged list:")
	l3.Print()

}
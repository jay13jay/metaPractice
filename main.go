package main

import (
	"fmt"

	"github.com/jay13jay/metaPractice/linkedlist"
)

func main() {
	list := &linkedlist.List{}

	for i := 0; i < 10; i++ {
		list.Append(i)
	}
	insertNode := list.Head
	for i := 0; i <= 4; i++ {
		insertNode = insertNode.Next
	}
	list.Insert(75, insertNode)

	list.Prepend(51)
	list.Delete(insertNode.Next)
	err := list.DeleteValue(100)
	if err != nil {
		fmt.Printf("Could not delete\nError: %v", err)
	}
	_ = list.DeleteValue(4)

	list.Print()
	list.Reverse()
	list.Print()
}
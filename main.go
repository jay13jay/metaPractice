package main

import (
	"github.com/jay13jay/metaPractice/mergesort"
	"github.com/jay13jay/metaPractice/quicksort"

	"github.com/jay13jay/metaPractice/utils"
)

var size int = 10000
var size2 int = 20000
func main() {
	// var arr = []int{10,7,8,9,1,5}
	// fmt.Println("QS:")
	utils.BruteFunc(quicksort.QuickSort, mergesort.MS, size)
}
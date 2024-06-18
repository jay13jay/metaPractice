package main

import (
	"github.com/jay13jay/metaPractice/mergesort"
	"github.com/jay13jay/metaPractice/utils"
)

var size int = 1000000
func main() {
	utils.BruteFunc(mergesort.MS, size)
}
package main

import (
	"fmt"

	"github.com/jay13jay/metaPractice/tunnel"
)

func main() {
	// var C int64 = 10
	// var N int32 = 2
	// var A []int64 = []int64{1, 6}
	// var B []int64 = []int64{3, 7}
	// var K int64 = 7
	var C int64 = 50
	var N int32 = 3
	var A []int64 = []int64{39, 19, 28}
	var B []int64 = []int64{49, 27, 35}
	var K int64 = 15
	ret := tunnel.Solve(C, N, A, B, K)
	fmt.Println("Time required: ", ret)
	
}
package main

import (
	"fmt"

	"github.com/jay13jay/metaPractice/hops"
)

func main() {
	var N int64 = 6 	// Number of pads
	var F int32 = 3 	// Number of frogs
	P := []int64{5, 2, 4} 	// Occupied pads
	
	ret := hops.SolveHops(N, F, P)
	fmt.Println(ret)
}
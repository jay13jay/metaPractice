package main

import (
	"fmt"

	"github.com/jay13jay/metaPractice/hops"
)

func main() {
	var N int64 = 6
	var F int32 = 3
	P := []int64{5, 2, 4}
	
	ret := hops.SolveHops(N, F, P)
	fmt.Println(ret)
}
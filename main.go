package main

import (
	"fmt"

	"github.com/jay13jay/metaPractice/rotarylock"
)

func main() {
	N := int32(10)
	M := int32(4)
	C :=[]int32{9, 4, 4, 8}

	time := rotarylock.Solve(N, M, C)
	fmt.Println(time)
}
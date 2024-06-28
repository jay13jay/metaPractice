package main

import (
	"fmt"

	"github.com/jay13jay/metaPractice/art"
)

func main() {
	n := int32(9)
	l := []int32{6, 3, 4, 5, 1, 6, 3, 3, 4}
	d := "ULDRULURD"

	ret := art.Solve(n, l, d)

	fmt.Println(ret)
}
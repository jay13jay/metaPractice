package main

import (
	"fmt"

	"github.com/jay13jay/metaPractice/spiral"
)

func main() {
	matrix := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	ret := spiral.Solve(matrix)

	fmt.Printf("Spiraled Matrix:\n%v\n", ret)
}
package main

import (
	"fmt"

	"github.com/jay13jay/metaPractice/scoreboard"
)

func main() {
	N := int32(6)
	var S = []int32{1, 2, 3, 4, 5, 6}


	scores := scoreboard.Solve(N, S)	
	fmt.Println(scores)
}
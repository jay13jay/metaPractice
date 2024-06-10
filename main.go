package main

import (
	"fmt"

	"github.com/jay13jay/metaPractice/battleship"
)


func main() {
	R := int32(5)
	C := int32(3)
	G := [][]int32{
			{1, 1, 0},
			{0, 0, 0},
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		}
	fmt.Println(battleship.GetHitProbability(R, C, G))
}
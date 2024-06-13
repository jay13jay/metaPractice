package main

import (
	"fmt"

	"github.com/jay13jay/metaPractice/fib"
)


func main() {
	// var N, K, M int64
	// N = 15 	// 
	// K = 2	// free space requirement
	// M = 3		// num of currently seated
	// // S := []int64{2, 6}
	// S := []int64{11,6,14}
	// // [0] [x][3][4][5][6][7][8][9][10][x][12][13][x][15]

	// res := cafeteria.GetMaxAdditionalDinersCount(N, K, M, S)

	// fmt.Printf("There are %d seats available\n", res)
	n := 10
	res := fib.CalcFib(n)
	fmt.Printf("Fibonacci of %d is %d\n", n, res)
}
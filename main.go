package main

import (
	"fmt"
	"time"

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
	
	// n := 20
	for n := 10; n <= 1000000; n = n + n {
		start := time.Now()
		res := fib.CalcFib(n)
		end := time.Now()
		timeDiff := end.Sub(start)
		fmt.Printf("%d\t%v\t%v\n", n, timeDiff, fib.FormatScientific(res))
	}
}
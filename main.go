package main

import (
	"fmt"
	"time"

	"github.com/jay13jay/metaPractice/gridtraveler"
)

const end = 10000000000
const start = end - 3
func main() {
	// fmt.Println("DP:")
	// for i := start; i <= end; i+=2 {
	// 	start := time.Now()
	// 	// fmt.Printf("Grid: %dX%d\tPaths: %d\n", i, i, gridtraveler.UniquePaths(i, i) )
	// 	_ = gridtraveler.UniquePaths(i, i)
	// 	end := time.Now()
	// 	diff := end.Sub(start)
	// 	fmt.Printf("Paths: %d\tTime: %s\n", i, diff)
	// }
	fmt.Println("\nCombinatory:")
	for i := start; i <= end; i++ {
		start := time.Now()
		// fmt.Printf("Grid: %dX%d\tPaths: %d\n", i, i, gridtraveler.UniquePathsCombinatory(i, i) )
		_ = gridtraveler.UniquePathsCombinatory(i, i)
		end := time.Now()
		diff := end.Sub(start)
		fmt.Printf("Grid:%d X %d\tTime: %s\n", i, i, diff)
	}
}
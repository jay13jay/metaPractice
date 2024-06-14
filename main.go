package main

import (
	"fmt"
	"time"

	"github.com/jay13jay/metaPractice/gridtraveler"
)

func main() {
	for i := 1000; i <= 10000; i+=2 {
		start := time.Now()
		fmt.Printf("Grid: %dX%d\tPaths: %d\n", i, i, gridtraveler.UniquePaths(i, i) )
		end := time.Now()
		diff := end.Sub(start)
		fmt.Printf("Time taken: %s\n", diff)
	}
}
package utils

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type fn func(arr []int)

func BruteFunc(f1, f2 fn, size int) {
	w := tabwriter.NewWriter(os.Stdout, 2, 4, 2, ' ', 0)
	fmt.Fprintln(w, "Arr size\tQuick Sort Time\t\tMerge Sort Time")
	for t := 0; t < 10; t++ {
		arr := CreateSlice(size)
		start1 := time.Now()
		f1(arr)
		end1 := time.Now()
		tDiff1 := end1.Sub(start1)

		arr = CreateSlice(size) // Recreate the slice for the second function to ensure fairness
		start2 := time.Now()
		f2(arr)
		end2 := time.Now()
		tDiff2 := end2.Sub(start2)

		var arrow string
		if tDiff1 > tDiff2 {
			arrow = "->"
		} else {
			arrow = "<-"
		}
		fmt.Fprintf(w, "Arr size: %d\t%s\t%s\t%s\n", size, tDiff1, arrow, tDiff2)
		// size += size / 2
		size += 1000
	}
	w.Flush()
}

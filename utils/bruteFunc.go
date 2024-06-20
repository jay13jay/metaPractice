package utils

import (
	"fmt"
	"image/color"
	"os"
	"text/tabwriter"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type fn func(arr []int)

func BruteFunc(f1, f2 fn, size int) {
	w := tabwriter.NewWriter(os.Stdout, 2, 4, 2, ' ', 0)
	fmt.Fprintln(w, "Arr size\tQuick Sort Time\t\tMerge Sort Time")

	var sizes []int
	var quickSortTimes []time.Duration
	var mergeSortTimes []time.Duration
	for t := 0; t < 10; t++ {
		arr := CreateSlice(size)
		arr2 := CreateSlice(size)
		start1 := time.Now()
		f1(arr)
		end1 := time.Now()
		tDiff1 := end1.Sub(start1)

		start2 := time.Now()
		f2(arr2)
		end2 := time.Now()
		tDiff2 := end2.Sub(start2)

		var arrow string
		if tDiff1 > tDiff2 {
			arrow = "->"
		} else {
			arrow = "<-"
		}
		fmt.Fprintf(w, "Arr size: %d\t%s\t%s\t%s\n", size, tDiff1, arrow, tDiff2)

		sizes = append(sizes, size)
		quickSortTimes = append(quickSortTimes, tDiff1)
		mergeSortTimes = append(mergeSortTimes, tDiff2)

		// size += size / 2
		size *= 2
	}
	w.Flush()
	plotResults(sizes, quickSortTimes, mergeSortTimes)
}


func plotResults(sizes []int, quickSortTimes, mergeSortTimes []time.Duration) {
	p := plot.New()

	p.Title.Text = "Sorting Algorithms Performance"
	p.X.Label.Text = "Array Size"
	p.Y.Label.Text = "Time (ms)"

	quickSortPoints := make(plotter.XYs, len(sizes))
	mergeSortPoints := make(plotter.XYs, len(sizes))

	for i := range sizes {
		quickSortPoints[i].X = float64(sizes[i])
		quickSortPoints[i].Y = float64(quickSortTimes[i].Microseconds())
		mergeSortPoints[i].X = float64(sizes[i])
		mergeSortPoints[i].Y = float64(mergeSortTimes[i].Microseconds())
	}

	quickSortLine, err := plotter.NewLine(quickSortPoints)
	if err != nil {
		panic(err)
	}
	quickSortLine.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255} // Red color

	mergeSortLine, err := plotter.NewLine(mergeSortPoints)
	if err != nil {
		panic(err)
	}
	mergeSortLine.Color = color.RGBA{B: 255, A: 255} // Blue color

	p.Add(quickSortLine, mergeSortLine)
	p.Legend.Add("Quick Sort", quickSortLine)
	p.Legend.Add("Merge Sort", mergeSortLine)

	if err := p.Save(10*vg.Inch, 5*vg.Inch, "sort_performance.png"); err != nil {
		panic(err)
	}
}
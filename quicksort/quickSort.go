package quicksort

// find the pivot point to sort > and < elements to
// > elements to the left
// < elements to the right
func partition(arr []int, low, high int) int {
    // Choose the pivot as the median of the first, middle, and last elements
    mid := low + (high-low)/2
    pivot := median(arr[low], arr[mid], arr[high])

    // Partition the array
    i := low + 1
    for j := low + 1; j <= high; j++ {
        if arr[j] < pivot {
            arr[i], arr[j] = arr[j], arr[i]
            i++
        }
    }
    arr[low], arr[i-1] = arr[i-1], arr[low]

    return i - 1
}

// takes the median of 3 values to pick a good starting pivot
func median(a, b, c int) int {
    if a < b {
        if b < c {
            return b
        } else if a < c {
            return c
        } else {
            return a
        }
    } else {
        if a < c {
            return a
        } else if b < c {
            return c
        } else {
            return b
        }
    }
}

// Takes a slice of ints, and the low and high ints of the slice
func sort(arr []int, low, high int) {
	if low < high {
		// get pivot point from partition func
		pivot := partition(arr, low, high)

		sort(arr, low, pivot-1)  // sort left of pivot
		sort(arr, pivot+1, high) // sort right of pivot
	}
}

// Takes a slice of ints. This stub function was chosen
// in order to be able to pass in just the array, making it have a similar
// signature to the Merge Sort implementation
func QuickSort(arr []int) {
	high := len(arr) - 1
	low := 0
	sort(arr, low, high)
	// fmt.Printf("QS Slice: %v ", arr)
}

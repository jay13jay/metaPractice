package mergesort

// MS sorts the given slice in-place using the merge sort algorithm.
func MS(arr []int) {
	mergeSortInPlaceHelper(arr, 0, len(arr)-1)
	// fmt.Printf("MS Slice: %v ", arr)
}

func mergeSortInPlaceHelper(arr []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		mergeSortInPlaceHelper(arr, left, mid)
		mergeSortInPlaceHelper(arr, mid+1, right)
		mergeInPlace(arr, left, mid, right)
	}
}

func mergeInPlace(arr []int, left, mid, right int) {
	// Create a temporary slice to store the merged elements
	temp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0

	// Merge the two sorted subarrays into the temporary slice
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
	}

	// Copy the remaining elements from the left subarray, if any
	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
	}

	// Copy the remaining elements from the right subarray, if any
	for j <= right {
		temp[k] = arr[j]
		j++
		k++
	}

	// Copy the merged elements back to the original array
	copy(arr[left:], temp)
}

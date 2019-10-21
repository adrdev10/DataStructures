package main

import "fmt"

func main() {
	var arr = [9]int{10, 20, 5, 9, 100, 30, 1, 0, 99}
	quickSort(&arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr *[9]int, lo, hi int) {
	if lo < hi {
		pivot := partition(arr, lo, hi)
		quickSort(arr, lo, pivot-1)
		quickSort(arr, pivot+1, hi)
	}
}

// Lomuto partition scheme
func partition(arr *[9]int, lo, hi int) int {
	pivot := arr[hi] //99
	i := lo
	for j := lo; j < hi; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[hi] = arr[hi], arr[i]
	return i
}

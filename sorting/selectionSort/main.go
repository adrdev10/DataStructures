package main

import "fmt"

func main() {
	var arr = [10]int{10, 20, 5, 9, 100, 30, 1, 0, 1999}
	selectionSorter(&arr)
	fmt.Println(arr)
}

func selectionSorter(arr *[10]int) {
	for i := 0; i < len(arr)-1; i++ {
		var min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[min], arr[i] = arr[i], arr[min]
	}
}

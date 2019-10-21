package main

import "fmt"

func main() {
	var arr = [12]int{10, 20, 5, 9, 100, 30, 1, 0, 1999, 200000, 10000, -1}
	insertionSorter(&arr)
	fmt.Println(arr)
}

func insertionSorter(arr *[12]int) {
	for i := 1; i < len(*arr); i++ {
		j := i
		for j > 0 {
			if arr[j] < arr[j-1] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
			j--
		}
	}
}

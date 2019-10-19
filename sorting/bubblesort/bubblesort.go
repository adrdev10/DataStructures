package main

import "fmt"

func main() {
	var arr = [10]int{10, 20, 5, 9, 100, 30, 1, 0, 1999}
	bubbleSorter(&arr)
	fmt.Println(arr)
}

func bubbleSorter(arr *[10]int) {
	var num int
	num = 10
	var isSwapped bool
	isSwapped = true
	for isSwapped {
		isSwapped = false
		var i int
		for i = 1; i < num; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				isSwapped = true
			}
		}
	}
}

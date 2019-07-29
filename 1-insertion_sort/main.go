package main

import (
	"fmt"
)

func InsertionSort(arr []int) {
	fmt.Print("\n===========================================\n")
	for i, v := range arr {
		j := i - 1
		for j >= 0 && arr[j] > v {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = v
		fmt.Printf("Array: %v\n", arr)
	}

	fmt.Print("===========================================\n\n")
}

func main() {
	unsorted := []int{100, 6, 34, 28, 97, 23, 61, 2, 7, 1, 44, 0}
	InsertionSort(unsorted)
	fmt.Printf("Result: %v", unsorted)
}

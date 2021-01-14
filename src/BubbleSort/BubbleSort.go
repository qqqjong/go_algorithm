package main

import "fmt"

func bubbleSort(arr []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
	}
	return arr
}

func main() {
	fmt.Println(bubbleSort([]int{1, 2, 3, 4, 7, 6}))
}

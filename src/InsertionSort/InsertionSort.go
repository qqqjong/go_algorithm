package main

import "fmt"

func InsertionSort(arr []int) []int {
	for currentIndex := 1; currentIndex < len(arr); currentIndex++ {
		temporary := arr[currentIndex]
		iterator := currentIndex
		for ; iterator > 0 && arr[iterator-1] >= temporary; iterator-- {
			arr[iterator] = arr[iterator-1]
		}
		arr[iterator] = temporary
	}
	return arr
}

func main() {
	fmt.Println(InsertionSort([]int{6,1,2,7,5,4}))
}
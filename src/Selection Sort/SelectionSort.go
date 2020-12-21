package main

import "fmt"

func selectionSort(arr []int) {
	len := len(arr)
	for i:=0; i<len-1; i++ {
		minIndex := i
		for j:=i; j<len; j++ {
			if arr[j] < arr[minIndex] {
				arr[j], arr[minIndex] = arr[minIndex], arr[j]
			}
		}
	}
	fmt.Println("\nAfter Selection Sort")
	for _, val := range arr {
		fmt.Println(val)
	}
}

func main() {
	sample := []int{3,4,5,2,1,10,5,999,12,1}
	selectionSort(sample)
}
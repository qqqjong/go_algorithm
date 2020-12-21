package main

import "fmt"

func selectionSort(arr []int) {
	len := len(arr)
	for i:=0; i<len-1; i++ {
		minIndex := i
		for j:=i+1; j<len; j++ {
			if arr[j] < arr[minIndex] {
				arr[j], arr[minIndex] = arr[minIndex], arr[j]
			}
		}
	}
	fmt.Println("\n선택정렬 완료")
	for _, val := range arr {
		fmt.Println(val)
	}
}

func main() {
	sample := []int{0,5,2,4,1,1,500,24}
	selectionSort(sample)
}
package main

import "fmt"

func bubbleSort(arr []int) {
	len := len(arr)
	for i:=0; i<len-1; i++ {
		for j:=0; j<len-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("\n버블정렬 완료")
	for _, val := range arr {
		fmt.Println(val)
	}
}

func main() {
	sample := []int{3,1,2,3,6,7,10,50,11}
	bubbleSort(sample)
}
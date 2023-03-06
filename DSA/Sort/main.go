package main

import (
	"datasort/bubble"
	"fmt"
)

func main() {

	fmt.Println("Demo: Bubble sort...")

	dataList := []int{4, 3, 8, 11, 1, 6, 21, 12, 15, 35, 26, 78, 54, 36, 77}
	fmt.Println("Before Sort:", dataList)
	dataList = bubble.BubbleSort(dataList)
	fmt.Println("Sorted List is: ", dataList)
}

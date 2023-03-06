package main

import (
	"fmt"
	binarysearch "search/BinarySearch"
	linearsearch "search/LinearSearch"
)

func main() {

	fmt.Println("Demo: Linear Search....")

	dataList := []int{3, 4, 5, 6, 7, 8, 21, 33, 44, 79, 86}

	pos, found := linearsearch.LinearSearch(dataList, 21)

	if found {
		fmt.Println("Found element at position:", pos)
	} else {
		fmt.Println("Not found")
	}

	stringList := []string{"Hello!", " and", " welcome", " to", " linear", " search!!"}

	pos, found = linearsearch.LinearSearch(stringList, 20)

	if found {
		fmt.Println("Found element at:", pos)
	} else {
		fmt.Println("Not found")
	}

	pos, found = linearsearch.LinearSearch(stringList, " welcome")

	if found {
		fmt.Println("Found element at:", pos)
	} else {
		fmt.Println("Not found")
	}

	pos, found = binarysearch.BinarySearch(dataList, 44)

	if found {
		fmt.Println("Found element at:", pos)
	} else {
		fmt.Println("Not found")
	}

	pos, found = binarysearch.BinarySearch(dataList, 56)

	if found {
		fmt.Println("Found element at:", pos)
	} else {
		fmt.Println("Not found")
	}

}

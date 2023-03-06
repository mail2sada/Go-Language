package main

import (
	"fmt"
	"linkedlist/single"
)

func main() {
	list := single.CreateList()

	list.AddItem(10)
	list.AddItem(20)
	list.AddItem("Hello")
	list.AddItem(26.22)

	elements := list.GetList()

	for _, val := range elements {
		fmt.Println(val)
	}

}

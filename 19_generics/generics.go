package main

import (
	"fmt"
)

func PrintSlice[T any](items []T) { //if you want specific types to allow use T int | string
	for _, item := range items {
		fmt.Println(item)
	}
}

//	func PrintStringSlice(items []string) {
//		for _, item := range items {
//			fmt.Println(item)
//		}
//	}
type stack[T any] struct {
	elements []T
}

func main() {
	// nums:= []int{1,2,3}

	// names := []string{"golang", "java"}
	// // PrintStringSlice(names)
	// PrintSlice(names)
	// PrintSlice(nums)
	//LIFO
myStack := stack[string]{
	elements: []string{"go", "java"},
}
fmt.Println(myStack)
}

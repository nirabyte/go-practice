package main

import "fmt"

func PrintSlice(items []int) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {
	PrintSlice([]int{1,2,3})
}

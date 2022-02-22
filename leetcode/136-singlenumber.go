package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{2, 5, 2}
	result := singleNumber(a)

	fmt.Println("Result-->", result)
}

func singleNumber(arr []int) int {
	sort.Ints(arr)
	lastEle := len(arr) - 1
	for i := 0; i < lastEle; i += 2 {
		if arr[i] != arr[i+1] {
			return arr[i]
		}
	}
	return arr[lastEle]
}

package main

import "fmt"

func main() {

	array1 := []int{1, 2, 3, 4, 5}
	target := 5
	a := searchInsert(array1, target)
	fmt.Println(a)
}

func searchInsert(nums []int, target int) int {
	var i int
	for i = 0; i < len(nums); i++ {
		if nums[i] == target {
			break
		}
	}
	return i
}

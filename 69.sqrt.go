package main

import (
	"fmt"
)

func main() {
	a := 10
	result := sqrt(a)
	fmt.Println(result)
}

func sqrt(a int) int {
	var mid int
	var sqr int

	start, end := 0, a
	for start < end {
		mid = (start + end) / 2
		fmt.Println(start, end)
		fmt.Println("mid", mid)
		if mid == start {
			fmt.Println("print")
			return mid
		}
		sqr = mid * mid
		fmt.Println("sqr", sqr)
		if a == sqr {
			return mid
		}
		if a > sqr {
			start = mid
		}
		if sqr > a {
			end = mid
		}
	}
	return 0
}

// func search(nums []int, target int) int {
// 	low, high := 0, len(nums)-1
// 	var mid int
// 	for low <= high {
// 		mid = (low + high) / 2
// 		if nums[mid] < target {
// 			low = mid + 1
// 		} else if nums[mid] > target {
// 			high = mid - 1
// 		} else {
// 			return mid
// 		}
// 	}

// 	return -1

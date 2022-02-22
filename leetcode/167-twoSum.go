package main

import "fmt"

func main() {
	arr := []int{2, 7, 11, 15}
	tar := 9

	result := twoSum(arr, tar)
	fmt.Println("Result:-", result)
}

func twoSum(arr []int, tar int) []int {
	end := len(arr) - 1
	start := 0
	for start < end {
		sum := arr[start] + arr[end]

		if sum < tar {
			start++
		} else if sum > tar {
			end--
		} else {
			break
		}
	}
	return []int{start + 1, end + 1}
}

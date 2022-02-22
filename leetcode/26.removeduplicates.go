package main

import "fmt"

func main() {

	arr := []int{1, 1, 1, 3, 4}
	a := removeDuplicate(arr)
	println(a)
}

func removeDuplicate(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}
	fmt.Println("array", nums[:j])
	return j
}

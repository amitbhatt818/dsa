package main

import (
	"fmt"
)

func main() {

	arr := []int{9, 2, 1, 1, 3, 9, 1}
	k := 2
	a := containsDuplicatetwo(arr, k)
	println(a)
}

func containsDuplicatetwo(nums []int, k int) bool {
	// records index of last seen for n.
	seen := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		idx, ok := seen[nums[i]]
		fmt.Println("map", seen)
		fmt.Println("okkkkkkkkk", ok)
		fmt.Println("ennnn", seen[nums[i]])
		fmt.Println("indxx", idx)
		if ok && i-idx <= k {
			return true
		}
		fmt.Println("i", i)
		seen[nums[i]] = i
	}

	return false
}

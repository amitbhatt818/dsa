package main

import "fmt"

func main() {

	arr := []int{8, 3, 2, 10, 8, 4}
	target := 6
	i, j := twoSum(arr, target)
	fmt.Println("Output1", i)
	fmt.Println("Output2", j)

}

func twoSum(arr []int, target int) (int, int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j]+arr[i] == target {
				return i, j
			}
		}
	}
	return 0, 0
}

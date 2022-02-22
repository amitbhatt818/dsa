package main

import (
	"fmt"
)

func main() {
	a := []int{1, -2, 3, -1, -2}

	result := subarraySum(a)
	fmt.Println(result)
}

func subarraySum(a []int) int {
	sum := a[0]
	max := sum

	for i := 1; i < len(a); i++ {
		sum = sum + a[i]
		if sum < a[i] {
			sum = a[i]
			fmt.Println("sum", sum, i)
		}
		if sum > max {
			max = sum
			fmt.Println("max", max)
		}
	}
	return max
}

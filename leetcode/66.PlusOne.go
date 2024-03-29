package main

import "fmt"

func main() {
	a := []int{1, 2, 9}

	result := plusOne(a)
	fmt.Println(result)
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] += 1
			return digits
		}
	}
	digits = append([]int{1}, digits...)
	return digits
}

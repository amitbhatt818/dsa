package main

import "fmt"

func main() {
	a := []int{1, 1, 2, 3, 2, 2}

	result := majorityEle(a)
	fmt.Println(result)
}

func majorityEle(arr []int) int {
	count := 0
	var candidate int

	for _, ele := range arr {
		if count == 0 {
			candidate = ele
		}
		if candidate == ele {
			count++
		} else {
			count--
		}
	}
	return candidate
}

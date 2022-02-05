package main

import "fmt"

func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(arr)
	result := maxArea(arr)
	fmt.Println("result:", result)
}
func maxArea(height []int) int {

	l := len(height) - 1
	max := 0
	maxAreaofContainer := 0
	i := 0
	for i < l {
		if height[i] < height[l] {
			max = (l - i) * height[i]
			i++
		} else {
			max = (l - i) * height[l]
			l--
		}
		if max > maxAreaofContainer {
			maxAreaofContainer = max
		}
	}
	return maxAreaofContainer
}

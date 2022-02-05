package main

import "fmt"

func main() {
	var a, b []int
	a = []int{1, 2, 3, 4, 5}
	b = []int{5, 7, 8, 9, 10, 11}
	len1 := len(a)
	len2 := len(b)
	len3 := len1 + len2
	c := make([]int, len3)
	result := merge(a, b, c, len1, len2)
	fmt.Println("Sorted Array:-->", result)

}

func merge(a []int, b []int, c []int, len1 int, len2 int) []int {
	var i, j, k int

	for i < len1 && j < len2 {
		if a[i] < b[j] {
			c[k] = a[i]
			k++
			i++
		} else {
			c[k] = b[j]
			k++
			j++

		}
	}
	for i < len1 {
		c[k] = a[i]
		i++
		k++
	}
	for j < len2 {
		c[k] = b[j]
		j++
		k++
	}
	return c
}

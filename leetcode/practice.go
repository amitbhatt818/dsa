package main

import "fmt"

func main() {
	n := 3
	result := noOfRows(n)
	fmt.Println(result)
}

func noOfRows(n int) [][]int {
	arr := [][]int{}
	if n >= 1 {
		arr = append(arr, []int{1})
	}

	if n >= 2 {
		arr = append(arr, []int{1, 1})
	}

	for i := 2; i < n; i++ {
		prev := arr[i-1]
		fmt.Print(prev)
		arr = append(arr, newRow(prev))

	}
	return arr
}

func newRow(a []int) []int {
	new := []int{1}
	fmt.Println(len(a))
	for i := 1; i < len(a); i++ {
		nextele := a[i] + a[i-1]
		new = append(new, nextele)
	}
	new = append(new, 1)
	return new
}

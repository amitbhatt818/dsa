package main

import "fmt"

func main() {
	a := 4
	result := climbStairs(a)
	fmt.Println(result)
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}
	next := 0
	first := 1
	second := 2
	for i := 2; i < n; i++ {
		next = first + second
		first = second
		second = next
	}
	return next
}

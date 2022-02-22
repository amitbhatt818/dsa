package main

import (
	"fmt"
)

func main() {
	n := 10
	result := trailingZero(n)
	fmt.Println(result)
}

func trailingZero(n int) int {
	var r int
	i := 5
	for i <= n {
		r = n/i + r
		fmt.Println("R", r)
		i = i * 5
	}
	return r
}

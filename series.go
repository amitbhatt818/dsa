package main

import "fmt"

func main() {
	a := 1
	b := 3
	c := 6
	N := 100
	elementseries(a, b, c, N)
}

func elementseries(a, b, c, N int) {
	diff1 := b - a
	diff2 := c - b
	result := a

	for i := 1; i < N; i++ {
		if i%2 == 0 {
			result += diff2
		} else {
			result += diff1
		}
	}
	fmt.Println(result)
}

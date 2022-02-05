package main

import "fmt"

func main() {
	s := "helloworld"
	result := lengthLast(s)
	fmt.Println(result)
}

func lengthLast(s string) int {
	count := 0
	for i := len(s) - 1; i > -1; i-- {
		if string(s[i]) != " " {
			count++
		} else if count > 0 {
			println("else-if")
			return count
		}
	}
	return count
}

package main

import "fmt"

func main() {
	var result int
	num := 5
	c := make(chan int)
	go fact(num, c)
	result = <-c
	fmt.Println(result)
}

func fact(num int, c chan int) {
	result := 1
	for i := 1; i <= num; i++ {
		result = result * i
	}

	c <- result
}

package main

import "fmt"

func main() {
	var length, res, car int = 1, 0, 0
	n := 5

	var arr [10]int
	//arr := make([]int, 3)
	arr[0] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < length; j++ {
			res = i*arr[j] + car
			arr[j] = res % 10
			car = res / 10
		}
		for car != 0 {
			arr[length] = car
			length++
			car = car / 10
		}
	}
	for i := length - 1; i >= 0; i-- {
		fmt.Print(arr[i])
	}
	//fmt.Println(arr)
}

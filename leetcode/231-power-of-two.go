package main

import "fmt"

func main() {
	var a int
	a = 19
	result := sumOfTwo(a)
	fmt.Println("result:", result)

}

func sumOfTwo(a int) bool {

	if a == 1 {
		return true
	}
	if a%2 != 0 || a == 0 {
		return false
	}
	return sumOfTwo(a / 2)
}

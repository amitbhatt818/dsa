package main

import "fmt"

func main() {
	a := "1110"
	b := "1"
	//
	result := addBinary(a, b)
	fmt.Println("-------->", result)
}

func addBinary(a string, b string) string {
	var p1 int = len(a) - 1
	var p2 int = len(b) - 1
	var result string
	var carry int

	for p1 >= 0 || p2 >= 0 {
		var i1 int
		var i2 int
		var sum int = carry
		if p1 >= 0 {
			i1 = int(a[p1] - '0')
			fmt.Println("P1", i1)
		}
		if p2 >= 0 {
			i2 = int(b[p2] - '0')
			fmt.Println("p2", i2)
		}
		sum += i1 + i2
		fmt.Println("start", result)
		fmt.Println("sum", sum)

		switch sum {
		case 0:
			result = "0" + result
			fmt.Println("result", result)
			carry = 0
		case 1:
			result = "1" + result
			carry = 0
			fmt.Println(result)
		case 2:
			result = "0" + result
			carry = 1
			fmt.Println(result)
		case 3:
			result = "1" + result
			carry = 1
			fmt.Println(result)
		}
		p1 -= 1
		p2 -= 1
		fmt.Println("loop", result)
	}

	if carry == 1 {
		result = "1" + result
	}

	return result
}

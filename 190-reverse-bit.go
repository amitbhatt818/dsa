package main

import (
	"fmt"
	"strconv"
)

func main() {

	//var n uint64
	n := "00000010100101000001111010011100"
	bits_number := 32
	number, _ := strconv.ParseUint(n, 2, bits_number)
	fmt.Println("numberrr", number)
	result := reverseBits(uint32(number))
	fmt.Println(result)
}

func reverseBits(num uint32) uint32 {
	var result uint32

	for i := 0; i < 32; i++ {
		result *= 2
		fmt.Println("Result:", result)
		if num != 0 {
			result += num % 2
			fmt.Println("num%2:", num%2)
			fmt.Println("Result for lop-->", result)
			num /= 2
			fmt.Println("num:", num)
		}
	}

	return result
}

package main

import (
	"fmt"
	"strconv"
)

func main() {

	//var n uint64
	n := "1111111111111000001111010011101"
	bits_number := 32
	number, _ := strconv.ParseUint(n, 2, bits_number)
	result := hammingWeight(uint32(number))
	fmt.Println("re---->", result)
}

func hammingWeight(num uint32) uint32 {
	var count uint32
	for num != 0 {
		count += num & 1
		num >>= 1
	}
	return count
}

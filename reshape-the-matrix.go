package main

import "fmt"

func main() {
	arr := [][]int{{0, 1, 2, 3}, // row=2 c=4 arr
		{4, 5, 6, 7},
	}

	newr := 4
	newc := 2

	result := matrixReshape(arr, newr, newc)
	fmt.Println("Arr::", result)

}

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if canReshape(nums, r, c) {
		return reshape(nums, r, c)
	}
	return nums
}

func canReshape(nums [][]int, r, c int) bool {
	row := len(nums)
	colume := len(nums[0])
	if row*colume == r*c {
		return true
	}
	return false
}

func reshape(nums [][]int, r, c int) [][]int {
	newShape := make([][]int, r)
	for index := range newShape {

		newShape[index] = make([]int, c)

	}
	rowIndex, colIndex := 0, 0
	for _, row := range nums {
		fmt.Println("row inside loop", row)
		for _, col := range row {
			fmt.Println("colindex", colIndex)
			fmt.Println("rowindiex", rowIndex)
			if colIndex == c {
				colIndex = 0
				rowIndex++
			}
			newShape[rowIndex][colIndex] = col
			colIndex++
			fmt.Println("for loop end")
			fmt.Println("print arrayyyyyyy", newShape)
		}
	}
	return newShape
}

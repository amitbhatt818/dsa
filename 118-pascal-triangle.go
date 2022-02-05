package main

import "fmt"

func main() {
	num := 5
	result := generate(num)

	fmt.Println("------------", result[num-1])

}

func generate(num int) [][]int {

	result := [][]int{}
	if num >= 1 {
		result = append(result, []int{1})

	}
	if num >= 2 {
		result = append(result, []int{1, 1})
	}
	fmt.Println("...", result)
	for i := 2; i < num; i++ {
		prev := result[i-1]
		fmt.Println(prev)
		result = append(result, createRow(prev))
		fmt.Println("new result", result)

	}

	return result

}

func createRow(prev []int) []int {
	newRow := []int{1}
	for i := 1; i < len(prev); i++ {
		nextEle := prev[i] + prev[i-1]
		newRow = append(newRow, nextEle)
	}
	newRow = append(newRow, 1)
	fmt.Println("new row", newRow)
	return newRow
}

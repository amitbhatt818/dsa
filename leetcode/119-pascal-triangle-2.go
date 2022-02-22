package main

import "fmt"

func main() {
	num := 4
	result := createRows(num)

	fmt.Println("------------", result)

}

func createRows(num int) []int {
	newRow := []int{1}
	for i := 0; i < num; i++ {
		for j := len(newRow) - 1; j > 0; j-- {
			newRow[j] = newRow[j-1] + newRow[j]
			fmt.Println("newrowwww:", newRow[j])
		}

		newRow = append(newRow, 1)
		fmt.Println("newrowwwwafter append:", newRow)
	}
	return newRow
}

// func createRow(prev []int) []int {
// 	newRow := []int{1}
// 	for i := 1; i < len(prev); i++ {
// 		nextEle := prev[i] + prev[i-1]
// 		newRow = append(newRow, nextEle)
// 	}
// 	newRow = append(newRow, 1)
// 	fmt.Println("new row", newRow)
// 	return newRow
// }

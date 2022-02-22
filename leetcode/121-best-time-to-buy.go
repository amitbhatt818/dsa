package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}

	result := maxProfit(prices)
	fmt.Println(result)
}

// ----Methord 1 -----
// func maxProfit(prices []int) int {
// 	var index int
// 	min := prices[0]
// 	for i := 1; i < len(prices); i++ {
// 		if prices[i] < min {
// 			min = prices[i]

// 			index = i
// 		}
// 		// fmt.Println("smallest element", min, i)
// 	}
// 	max := prices[index]
// 	for j := index + 1; j < len(prices); j++ {
// 		if prices[j] > max {
// 			max = prices[j]
// 		}
// 	}
// 	bestbuy := max - min

// 	return bestbuy
// }

// ------methord 2

func maxProfit(prices []int) int {
	maxdiff := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			diff := prices[i] - min
			if maxdiff < diff {
				maxdiff = diff
			}
		}
	}
	return maxdiff
}

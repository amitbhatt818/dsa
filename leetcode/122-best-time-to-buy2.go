package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}

	result := maxProfit(prices)
	fmt.Println(result)
}

func maxProfit(prices []int) int {
	profit := 0
	l := len(prices)

	for i := 0; i < l-1; i++ {
		if prices[i+1] > prices[i] {
			profit = profit + prices[i+1] - prices[i]
			fmt.Println("profit:", profit)
		}
	}

	return profit
}

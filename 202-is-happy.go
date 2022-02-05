package main

import "fmt"

func main() {
	n := 19

	result := isHappy(n)
	fmt.Println(result)
}
func isHappy(n int) bool {
	slow := findSquareSum(n)
	fmt.Println("slow:", slow)
	fast := findSquareSum(findSquareSum(n))
	fmt.Println("fast:", fast)

	for slow != fast {
		slow = findSquareSum(slow)
		fmt.Println("slow:Inside loop", slow)
		fast = findSquareSum(findSquareSum(fast))
		fmt.Println("fast:Inside loop", fast)
	}

	return slow == 1
}

func findSquareSum(num int) int {
	sum := 0
	digit := 0

	for num > 0 {
		digit = num % 10
		sum += digit * digit
		num /= 10
	}

	return sum
}

package main

import "fmt"

func main() {
	n := 7

	result := countPrimes(n)
	fmt.Println(result)
}

//rule out multiples
func countPrimes(n int) int {
	//true if not a prime, false if it is a prime
	set := make([]bool, n)
	var count int
	for i := 2; i < n; i++ {
		if !set[i] {
			for j := i; j < n; j += i {
				set[j] = true
			}
			count++
		}
	}
	return count
}

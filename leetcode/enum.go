// package main

// import "fmt"

// // Direction - Custom type to hold value for week day ranging from 1-4
// type Direction int

// // Declare related constants for each direction starting with index 1
// const (
// 	North Direction = iota + 1 // EnumIndex = 1
// 	East            = iota * 2 // EnumIndex = 2
// 	South           = iota     // EnumIndex = 3
// 	West            = iota     // EnumIndex = 4
// )

// func main() {
// 	var d Direction
// 	d = South
// 	fmt.Println(d) // Print : West
// 	fmt.Scanf("%s", &text1)
// 	arr := make([]string, 5)
// 	for i := range arr {
// 		fmt.Scanf("%s", &arr[i])
// 	}
// 	arr = append(arr, "tet", "pet")

// 	fmt.Println(arr)
// }

package main

import "fmt"

func main() {
	var text1, text2 string
	fmt.Scanf("%s", &text1)
	fmt.Scanf("%s", &text2)
	c := longestCommonSubsequence(text1, text2)
	fmt.Println(c)
}

func longestCommonSubsequence(text1 string, text2 string) int {
	var count int
	for i := 0; i <= len(text1); i++ {
		for j := 0; j <= len(text2); j++ {
			if text1[i] == text2[j] {
				count++
			}
		}
	}
	return count
}

// package main

// import "fmt"

// func main() {
// 	n := "AB"
// 	result := titleToNumber(n)
// 	fmt.Println(result)
// }
// func titleToNumber(s string) int {
// 	col := 0

// 	for _, char := range s {
// 		col = col * 26
// 		col = col + int(char-'A') + 1
// 	}

// 	return col
// }

package main

import "fmt"

func main() {
	a := "AAA"
	result := convertCharToInt(a)
	fmt.Println(result)
}

func convertCharToInt(s string) int {
	r := 0
	for _, c := range s {
		r = r * 26
		r = r + int(c-'A') + 1

	}
	return r
}

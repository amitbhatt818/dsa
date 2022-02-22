// package main

// import "fmt"

// func main() {
// 	s := "VIIM"
// 	result := romanToInt(s)
// 	fmt.Println(result)

// }

// func romanToInt(s string) int {
// 	var result int
// 	m := map[int32]int{
// 		'I': 1,
// 		'V': 5,
// 		'X': 10,
// 		'L': 50,
// 		'C': 100,
// 		'D': 500,
// 		'M': 1000,
// 	}
// 	for index, item := range s {
// 		current := m[item] // 5
// 		// fmt.Println(current)
// 		var next int
// 		if index < len(s)-1 {
// 			//=
// 			next = m[int32(s[index+1])]
// 			fmt.Println("next value", next)
// 		}

// 		if current < next {
// 			result -= current
// 		} else {
// 			result += current
// 		}
// 	}

// 	return result
// }

package main

import "fmt"

func main() {
	s := "IV"
	result := romanToInt(s)
	fmt.Println(result)
	d := len(s)
	fmt.Println(d)
}

func romanToInt(s string) int {
	var next int
	var sum int
	m := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}
	for index, j := range s {
		current := m[j]
		if index < len(s)-1 {
			next = m[rune(s[index+1])]
		}
		if current < next {
			sum = sum - current
		} else {
			sum = sum + current
		}
	}
	return sum
}

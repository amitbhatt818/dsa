// package main

// import "fmt"

// func main() {
// 	n := 26
// 	result := convertToTitle(n)
// 	fmt.Println(result)
// }

// func convertToTitle(n int) string {
// 	m := map[int]string{
// 		1:  "A",
// 		2:  "B",
// 		3:  "C",
// 		4:  "D",
// 		5:  "E",
// 		6:  "F",
// 		7:  "G",
// 		8:  "H",
// 		9:  "I",
// 		10: "J",
// 		11: "K",
// 		12: "L",
// 		13: "M",
// 		14: "N",
// 		15: "O",
// 		16: "P",
// 		17: "Q",
// 		18: "R",
// 		19: "S",
// 		20: "T",
// 		21: "U",
// 		22: "V",
// 		23: "W",
// 		24: "X",
// 		25: "Y",
// 		26: "Z",
// 	}

// 	res := ""
// 	for n > 0 {
// 		d := n % 26
// 		fmt.Println("Print D", d)
// 		n = n / 26
// 		fmt.Println("print n", n)

// 		if d == 0 {
// 			d = 26
// 			fmt.Println("Indide loop d", d)
// 			n = n - 1
// 			fmt.Println("Indide loop n", n)
// 		}
// 		fmt.Println("pRINT M[D]", m[d])
// 		res = m[d] + res
// 		fmt.Println("Result", res)
// 	}
// 	return res
// }

package main

import "fmt"

func main() {
	a := 52
	result := convertIntToCHar(a)
	fmt.Println(result)
}

func convertIntToCHar(num int) string {
	result := ""
	m := map[int]string{
		1:  "A",
		2:  "B",
		4:  "D",
		5:  "E",
		6:  "F",
		7:  "G",
		8:  "H",
		9:  "I",
		10: "J",
		11: "K",
		12: "L",
		13: "M",
		14: "N",
		15: "O",
		16: "P",
		17: "Q",
		18: "R",
		19: "S",
		20: "T",
		21: "U",
		22: "V",
		23: "W",
		24: "X",
		25: "Y",
		26: "Z",
	}

	for num > 0 {
		mod := num % 26
		fmt.Println("Print D", mod)
		num = num / 26
		if mod == 0 {
			mod = 26
			num = num - 1
		}
		result = m[mod] + result
	}

	return result
}

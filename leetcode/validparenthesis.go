package main

import "fmt"

func main() {
	s := "{({{{{}}}})}"

	output := isValid(s)
	println(output)
}

func isValid(s string) bool {
	seen := map[rune]rune{
		// '(': ')',
		// '[': ']',
		// '{': '}',
		')': '(',
		']': '[',
		'}': '{',
	}

	stk := []rune{}

	for _, i := range s {
		if i == ')' || i == '}' || i == ']' {
			if len(stk) > 0 && seen[i] == stk[len(stk)-1] {
				stk = stk[:len(stk)-1]
				fmt.Printf("%c", stk)
				fmt.Println()
				continue
			}
		}
		stk = append(stk, i)
		fmt.Printf("%c", i)
		fmt.Println()
	}

	return len(stk) == 0
}

package main

import "fmt"

func main() {
	str := []string{"aww", "ae", "a"}
	a := longestSubstringSuffix(str)
	fmt.Println(a)
}

func longestSubstringSuffix(str []string) string {
	var res []rune
	strlen := len(str)
	if strlen == 0 {
		return ""
	}
	lenstring := len(str[0])
	for _, item := range str {
		lenItem := len(item)
		if lenItem < lenstring {
			lenstring = lenItem
			println(lenstring)
		}
	}
	for i := 0; i < lenstring; i++ {
		for j := 1; j < strlen; j++ {
			if str[0][i] != str[j][i] {
				return string(res)
			}
		}
		res = append(res, rune(str[0][i]))
	}
	return string(res)
}

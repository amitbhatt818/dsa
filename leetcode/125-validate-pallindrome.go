package main

import "fmt"

func main() {
	s := "A man, a plan, a canal: Panama"
	result := isPallindrome(s)

	fmt.Println(result)
}

func isPallindrome(s string) bool {
	if s == "" {
		return true
	}
	start := 0
	end := len(s) - 1
	for start < end {
		for start < end && !isAlphanumeric(s[start]) {
			start++
		}
		for start < end && !isAlphanumeric(s[end]) {
			end--
		}
		if start < end && isLowecase(s[start]) != isLowecase(s[end]) {
			return false
		}
		start++
		end--
	}
	return true
}

func isAlphanumeric(c byte) bool {
	if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		return true
	}
	return false
}

func isLowecase(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		c += 32
	}
	return c
}

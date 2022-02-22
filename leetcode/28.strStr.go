package main

func main() {

	str1 := "hello"
	str2 := "ll"
	a := strStr(str1, str2)
	println(a)
}
func strStr(haystack, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}

	for i := 0; i <= len(haystack)-n; i++ {
		if haystack[i:i+n] == needle {
			println(i)
			return i
		}
	}

	return -1
}

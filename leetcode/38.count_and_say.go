package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countAndSay2(n int) string {
	s := "1"
	for i := 0; i < n-1; i++ {
		lt, temp, count := s[0], "", 0
		for j := 0; j < len(s); j++ {
			if lt == s[j] {
				count += 1
			} else {
				temp = fmt.Sprintf("%s%d%c", temp, count, lt)
				lt = s[j]
				count = 1
			}
		}
		temp = fmt.Sprintf("%s%d%c", temp, count, lt)
		fmt.Println()
		fmt.Println("print temp", temp)
		s = temp
	}
	return s
}

func lss(n int) string {
	var result strings.Builder
	s1 := "1"
	if n <= 1 {
		return s1
	}
	temp := s1
	var j int
	for i := 2; i <= n; i++ {
		for j = 0; j < len(temp); j++ {
			count := 1
			ts := temp[j]
			for j < len(temp)-1 && temp[j] == temp[j+1] {
				count++
				j++
			}
			result.WriteString(strconv.Itoa(count) + string(ts))
		}
		temp = result.String()
		result.Reset()

	}
	return temp
}

func main() {

	s := countAndSay2(4)
	fmt.Println(s)

}

func countAndSay(n int) string {
	if n == 1 {
		fmt.Println("inside if condition", n)
		return "1"
	}
	fmt.Println("outside if condition", n)

	var ret []byte
	base := countAndSay(n - 1)
	fmt.Println("after function call", base)
	for i := 0; i < len(base); i++ {
		count := 1
		fmt.Println("inside for loop", len(base))
		b := base[i] // digits..
		j := i
		for j+1 < len(base) && b == base[j+1] {
			j++
			count++
		}
		kk := strconv.Itoa(count)
		ret = append(ret, []byte(kk)...)
		ret = append(ret, b)
		i = j
	}
	return string(ret)
}

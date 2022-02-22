package main

import "fmt"

func main() {
	s := "22"
	result := letterCombinations(s)
	fmt.Println(result)
}

func letterCombinations(digits string) []string {
	m := map[rune][]rune{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	fmt.Println(m)
	output := make([]string, 0)

	for _, d := range digits {
		tempSlice := make([]string, 0)
		for _, r := range m[d] {
			if len(output) == 0 {
				tempSlice = append(tempSlice, string(r))
			} else {
				for _, s := range output {
					tempSlice = append(tempSlice, s+string(r))
				}
			}
		}
		output = tempSlice
		// fmt.Println(output)
	}
	return output
}

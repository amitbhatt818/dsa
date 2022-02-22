package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 5, 6, 7, 9}
	result := summaryRanges(array)
	fmt.Println(result)
}
func summaryRanges(arr []int) []string {
	s := []string{}
	for i, j := 0, 0; j < len(arr); j++ {
		i = j
		fmt.Println("insude i")
		for j+1 < len(arr) && arr[j]+1 == arr[j+1] {
			j++
			fmt.Println("insude j")
		}
		if i == j {
			fmt.Println("insude equal")
			s = append(s, fmt.Sprintf("%v", arr[i]))
		} else {
			s = append(s, fmt.Sprintf("%v-->%v", arr[i], arr[j]))
		}
	}
	return s
}

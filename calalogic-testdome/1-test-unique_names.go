package main

import "fmt"

func uniqueNames(a, b []string) []string {

	for _, ele := range b {
		a = append(a, ele)
	}
	fmt.Println("print c:", a)
	unique := make(map[string]bool, len(a))
	us := make([]string, len(unique))
	for _, elem := range a {
		fmt.Println("unique:", !unique[elem])
		fmt.Println("unique without not:", unique[elem])
		if !unique[elem] {
			fmt.Println("unique---- inside:", !unique[elem])
			us = append(us, elem)
			unique[elem] = true
		}
	}
	return us
}

func main() {
	// should print Ava, Emma, Olivia, Sophia
	fmt.Println(uniqueNames(
		[]string{"Ava", "Emma", "Olivia"},
		[]string{"Olivia", "Sophia", "Emma"}))
}

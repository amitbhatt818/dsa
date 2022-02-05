package main

import "fmt"

func main() {
	integer := 132
	s := pall(integer)
	fmt.Println(s)

}

func pall(x int) string {
	old_val := x
	int_new := 0
	for x > 0 {
		remainder := x % 10
		int_new = remainder + int_new*10
		x = x / 10
	}
	if old_val == int_new {
		return "Pallindrome no."
	}
	return "not a Pallindrome no."
}

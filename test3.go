package main

import (
	"fmt"
)

type customer struct {
	name    string
	age     int
	address string
}

func main() {
	fmt.Println("Hello, playground")

	var cus customer = customer{"xyz", 1, "adv-def"}

	getName()
}
func getName() {
	fmt.Println(c.name)
}

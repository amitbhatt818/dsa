package main

import (
	"fmt"
	"strconv"
)

type Car interface {
	Color() string
	Speed() string
}

type Honda struct {
	ColorValue string
	SpeedValue int
}

type Hyundai struct {
	ColorValue string
	SpeedValue int
}

func (honda Honda) Color() string {
	return "Color of Honda Car is " + honda.ColorValue
}

func (honda Honda) Speed() string {
	return "Speed of Honda Car is " + strconv.Itoa(honda.SpeedValue)
}

func (hyundai Hyundai) Color() string {
	return "Color of Hyundai Car is " + hyundai.ColorValue
}

func (hyundai Hyundai) Speed() string {
	return "Speed of Hyundai Car is " + strconv.Itoa(hyundai.SpeedValue)
}

func main() {
	var car Car = Honda{"Red", 240}

	fmt.Println(car.Color())
	fmt.Println(car.Speed())

	fmt.Println("Value-", car)
	fmt.Printf("Type- %T", car)

	car = Hyundai{"White", 270}

	fmt.Println(car.Color())
	fmt.Println(car.Speed())

	fmt.Println("Value-", car)
	fmt.Printf("Type- %T", car)
}

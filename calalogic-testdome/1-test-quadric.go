package main

import (
	"fmt"
	"math"
)

func findRoots(a, b, c float64) (float64, float64) {
	var discriminant, root1, root2, imaginary float64
	discriminant = (b * b) - (4 * a * c)
	discriminant = (b * b) - (4 * a * c)

	if discriminant > 0 {
		root1 = (-b + math.Sqrt(discriminant)/(2*a))
		root2 = (-b - math.Sqrt(discriminant)/(2*a))
		return root1, root2
	} else if discriminant == 0 {
		root1 = -b / (2 * a)
		root2 = -b / (2 * a)
		return root1, root2
	} else if discriminant < 0 {
		root1 = -b / (2 * a)
		root2 = -b / (2 * a)
		imaginary = math.Sqrt(-discriminant) / (2 * a)
		return root1 + imaginary, root2 - imaginary
	}
	return 0, 0
}

func main() {
	x1, x2 := findRoots(2, 10, 8)
	fmt.Printf("Roots: %f, %f", x1, x2)
}

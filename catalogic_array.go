// All elements which are not in arr2 should be
//appendend to arr1 and vice versa for array2  also

package main

import "fmt"

func main() {

	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{3, 4, 5, 6}

	a, b := reverse(arr1, arr2)
	fmt.Println("array 1", a)
	fmt.Println("array 2", b)

}

func reverse(a []int, arr2 []int) ([]int, []int) {
	var isNotPresent bool
	var num int
	for _, ele := range a {

		for _, ele2 := range arr2 {
			if ele == ele2 {
				isNotPresent = true
			}
			if !isNotPresent {
				num = ele2
			}
		}
		if !isNotPresent {
			a = append(a, num)
		}
		isNotPresent = true

	}
	for _, ele := range arr2 {
		for _, ele2 := range a {
			if ele != ele2 {
				arr2 = append(arr2, ele2)
			}
		}

	}
	return a, arr2
}

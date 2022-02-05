// package main

// import "fmt"

// func main() {

// 	a := []int{1, 2, 2, 0}
// 	k := 2
// 	result := contains(a, k)
// 	fmt.Println(result)
// }
// func contains(array []int, k int) bool {
// 	m := make(map[int]int)
// 	for i, ele := range array {
// 		if m[i] == ele {
// 			return true
// 		}
// 		m[ele] = true
// 	}
// 	return false
// }

package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.)

}

package main

func main() {

	arr := []int{1, 2, 1, 4}
	a := containsDuplicate(arr)
	println(a)
}

func containsDuplicate(arr []int) bool {
	m := make(map[int]bool)

	for _, n := range arr {
		if m[n] == true {
			return true
		}
		m[n] = true
	}
	return false
}

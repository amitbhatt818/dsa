package main

import "fmt"

func main() {

	arr := []string{"a", "b", "c"}
	arr2 := []string{"a", "c", "c", "b"}
	a := checkMsnInEachNode(arr, arr2)
	println(a)
}

func checkMsnInEachNode(workerNodes []string, msnList []string) bool {
	workerNodesMap := make(map[string]int)
	mayastorNodeMap := make(map[string]int)

	for _, workerNodesMapElem := range workerNodes {
		workerNodesMap[workerNodesMapElem]++
	}
	fmt.Println("worker node map", workerNodesMap)
	for _, msnListElem := range msnList {
		mayastorNodeMap[msnListElem]++
	}
	fmt.Println("msn  map", mayastorNodeMap)
	for MapKey, MapVal := range workerNodesMap {
		fmt.Println("map----keyyy", MapKey)
		if mayastorNodeMap[MapKey] != MapVal {
			return false
		}
	}
	fmt.Println("cobined map  map", mayastorNodeMap)

	return true
}

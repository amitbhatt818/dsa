package main

import "fmt"

func main() {
	arr := [][]int{{1, 0, 0, 1}, // row=2 c=4 arr
		{1, 0, 0, 1},
	}
	fmt.Println(arr)

	result := numIslands1(arr)
	fmt.Println("Arr::", result)

}
func numIslands1(grid [][]int) int {
	fmt.Println("ooooo", len(grid))
	visited, res := make([][]bool, len(grid)), 0
	fmt.Println(visited)
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {

		for j := 0; j < len(grid[0]); j++ {

			if grid[i][j] == 1 && !visited[i][j] {
				numIslandsHelper(grid, i, j, visited)
				fmt.Println("inside if", i, j)
				res++
			}
		}
	}
	return res
}

func numIslandsHelper(grid [][]int, i, j int, visited [][]bool) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {

		return
	}
	if visited[i][j] || grid[i][j] == 0 {
		return
	}

	visited[i][j] = true
	numIslandsHelper(grid, i-1, j, visited)
	numIslandsHelper(grid, i, j-1, visited)
	numIslandsHelper(grid, i+1, j, visited)
	numIslandsHelper(grid, i, j+1, visited)
}

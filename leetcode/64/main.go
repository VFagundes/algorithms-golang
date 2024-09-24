package main

import "fmt"

func main() {
	a, b := [][]int{{1, 3, 1}, {1, 2, 1}, {4, 1, 1}}, [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}

	fmt.Println(minPathSum(a))
	fmt.Println(minPathSum(b))
}
func minPathSum(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	fmt.Println(grid)
	// Update the first column
	for i := 1; i < rows; i++ {
		grid[i][0] += grid[i-1][0]
	}

	// Update the first row
	for j := 1; j < cols; j++ {
		grid[0][j] += grid[0][j-1]
	}
	fmt.Println(grid)

	// Update the rest of the grid
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}

	return grid[rows-1][cols-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

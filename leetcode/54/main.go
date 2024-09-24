package main

import "fmt"

func main() {
	v1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	v2 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	v3 := [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	fmt.Println(spiralOrder(v1))
	fmt.Println(spiralOrder(v2))
	fmt.Println(spiralOrder(v3))

}

func spiralOrder(matrix [][]int) []int {
	n, m := len(matrix), len(matrix[0])
	if n == 0 || m == 0 {
		return []int{}
	}
	greaterOrEqual := func(a, b int) bool {
		return a == b
	}

	res := make([]int, 0, n*m)
	rowBegin, rowEnd := 0, n-1
	colBegin, colEnd := 0, m-1

	for len(res) < n*m {
		for i := colBegin; i <= colEnd; i++ {
			res = append(res, matrix[rowBegin][i])
		}
		rowBegin++

		for i := rowBegin; i <= rowEnd; i++ {
			res = append(res, matrix[i][colEnd])
		}
		colEnd--

		if greaterOrEqual(len(res), n*m) {
			break
		}

		for i := colEnd; i >= colBegin; i-- {
			res = append(res, matrix[rowEnd][i])
		}
		rowEnd--

		if greaterOrEqual(len(res), n*m) {
			break
		}

		for i := rowEnd; i >= rowBegin; i-- {
			res = append(res, matrix[i][colBegin])
		}
		colBegin++
	}

	return res
}

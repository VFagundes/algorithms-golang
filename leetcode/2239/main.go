package main

import (
	"fmt"
	"math"
)

func findClosestNumber(arr []int) int {
	r := arr[0]
	for _, v := range arr[1:] {
		abs1, abs2 := math.Abs(float64(r)), math.Abs(float64(v))
		if abs1 > abs2 || abs1 == abs2 && v > r {
			r = v
		}
	}
	return r
}
func main() {
	arr := []int{-4, -2, 1, 4, 8}
	fmt.Println(findClosestNumber(arr))
}

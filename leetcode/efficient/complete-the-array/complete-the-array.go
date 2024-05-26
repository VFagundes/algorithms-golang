package main

import (
	"fmt"
	"slices"
)

func main() {

	fmt.Printf("result: %d, expected = 2\n", completeArray([]int{4, 8, 6}))
	fmt.Printf("result: %d, expected = 4\n", completeArray([]int{-1, 3, 5}))
	fmt.Printf("result: %d, expected = 8\n", completeArray([]int{0, 5, 10}))
}
func completeArray(nums []int) int {
	result := 0
	slices.Sort(nums)
	min := nums[0]
	total := len(nums)
	max := nums[total-1] + 1
	result = max - (total + min)

	return result
}

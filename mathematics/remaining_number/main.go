package main

import (
	"fmt"
	"sort"
)

func findDistance(nums []int) int {
	//size max - min - len +1
	min, max := 0, 0
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	min = nums[0]
	max = nums[len(nums)-1]
	return max - min - len(nums) + 1

}

func main() {
	nums := [][]int{{3, 5, 7}, {2, 4, 70, 8}}
	for _, arr := range nums {
		fmt.Println(findDistance(arr))

	}

}

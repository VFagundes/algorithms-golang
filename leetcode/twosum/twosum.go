package main

import (
	"fmt"
)

// https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)
	for i, v := range nums {
		current, ok := m[v]
		if ok {
			return []int{current, i}
		}
		m[target-v] = i

	}

	return []int{}
}

func main() {
	arr := []int{2, 11, 15, 33, 7}
	fmt.Println(twoSum(arr, 9))
	arr2 := []int{2, 0, 33, 5, 7, 10}
	fmt.Println(twoSum(arr2, 10))
	arr3 := []int{2, 0, 23, 335, 300, 100}
	fmt.Println(twoSum(arr3, 323))

}

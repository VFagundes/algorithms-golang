package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {

	complementMap := map[int]int{}
	for i, v := range nums {
		pair := target - v
		if j, ok := complementMap[pair]; ok {
			return []int{j, i}
		}
		complementMap[v] = i
	}

	return []int{}
}

func main() {
	arr := []int{2, 7, 11, 15, 33}
	fmt.Println(twoSum(arr, 9))
	arr2 := []int{2, 0, 33, 5, 7, 10}
	fmt.Println(twoSum(arr2, 10))
	arr3 := []int{2, 0, 23, 335, 300, 100}
	fmt.Println(twoSum(arr3, 323))

}

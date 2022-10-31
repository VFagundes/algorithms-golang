package main

import (
	"fmt"
	"sort"
)

//https://leetcode.com/problems/3sum/
//[-1,0,1,2,-1,-4]
func threeSum(nums []int) [][]int {
	ret := [][]int{}
	sort.Ints(nums)
	cMap := map[int]int{}
	for i:=0

	return ret
}

func main() {
	nums:= []int{1,2,3}
	threeSum(nums)
	fmt.Println("voila")
}

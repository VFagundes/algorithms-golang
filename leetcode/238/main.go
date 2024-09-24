package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)
	pre, suf := 1, 1
	for i := 0; i < n; i++ {
		answer[i] = pre
		pre *= nums[i]
		//1,1,2,6
	}
	for i := n - 1; i >= 0; i-- {
		answer[i] *= suf
		suf *= nums[i]

	}
	return answer
}

func main() {

	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))

}

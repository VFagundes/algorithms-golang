package main

import "fmt"

func summaryRanges(nums []int) []string {
	answer := []string{}
	n := len(nums)
	i := 0
	for i < n {
		s := nums[i]
		for i < n-1 && nums[i+1] == nums[i]+1 {
			i++
		}
		if s == nums[i] {
			answer = append(answer, fmt.Sprintf("%d", s))
		} else {
			answer = append(answer, fmt.Sprintf("%d->%d", s, nums[i]))
		}
		i++
	}
	return answer
}

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}

package main

import "fmt"

var longestOnes = func(nums []int, k int) int {
	max, l, zeros := 0, 0, 0
	for r := 0; r < len(nums); r++ {
		if nums[r] == 0 {
			zeros++
		}
		for k < zeros {
			if nums[l] == 0 {
				zeros--
			}
			l++
		}
		if max < r-l+1 {
			max = r - l + 1
			continue
		}

	}

	return max
}

func main() {
	testCases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		{[]int{0, 0, 1, 1, 1, 0, 0}, 0, 3},
		{[]int{0, 0, 1, 1, 0, 0, 1, 1, 0, 1}, 1, 4},
		{[]int{1, 1, 1, 1, 1}, 3, 5},
		{[]int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1}, 2, 5},
	}

	for _, testCase := range testCases {
		result := longestOnes(testCase.nums, testCase.k)
		fmt.Printf("nums: %v, k: %d -> result: %d, expected: %d\n", testCase.nums, testCase.k, result, testCase.expected)
	}

}

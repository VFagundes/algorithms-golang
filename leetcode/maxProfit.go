package main

import (
	"fmt"
	"math"
)

//https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

func maxProfitSl1(prices []int) int {
	i, j := 0, 1
	max := 0
	for j < len(prices) {

		if prices[i] > prices[j] {
			i = j
		} else if max < prices[j]-prices[i] {
			max = prices[j] - prices[i]
		}
		j++

	}
	return max
}
func minF(a, b int) int {

	if a < b {
		return a
	}
	return b

}

func maxF(a, b int) int {

	if a > b {
		return a
	}
	return b

}

func maxProfitSl2(prices []int) int {

	max := 0
	min := int(math.MaxInt32)
	for _, v := range prices {
		min = minF(v, min)
		max = maxF(v-min, max)

	}
	return max
}

func main() {
	arr1 := []int{7, 1, 5, 3, 6, 4}
	arr2 := []int{7, 6, 4, 3, 1}
	arr3 := []int{7, 6, 5, 3, 6, 9}
	arr4 := []int{4, 5, 6, 7, 1, 10}
	fmt.Println(maxProfitSl2(arr1), " expected: 5")
	fmt.Println(maxProfitSl2(arr2), " expected: 0")
	fmt.Println(maxProfitSl2(arr3), " expected: 6")
	fmt.Println(maxProfitSl2(arr4), " expected: 9")
	fmt.Println("------------------------------------")
	fmt.Println(maxProfitSl1(arr1), " expected: 5")
	fmt.Println(maxProfitSl1(arr2), " expected: 0")
	fmt.Println(maxProfitSl1(arr3), " expected: 6")
	fmt.Println(maxProfitSl1(arr4), " expected: 9")
	fmt.Println("------------------------------------")

}

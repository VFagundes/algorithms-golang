package main

import "fmt"

func _max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxProfit(prices []int) int {
	index := prices[0]
	profit := 0
	// Input: prices = [7,1,5,3,6,4]
	for _, v := range prices[1:] {
		if index > v {
			index = v
		}
		profit = _max(v, index-v)
	}
	return profit

}
func main() {

	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

package main

import (
	"fmt"
	"sort"
)

var data []int

func binarySearch(arr []int, val, left, right int) []int {

	mid := (left + right) / 2
	if left > right || mid > len(arr)-1 {
		return []int{}
	}
	if arr[mid] == val {
		return []int{arr[mid]}
	} else if val < arr[mid] {
		return binarySearch(arr, val, left, mid-1)
	} else {
		return binarySearch(arr, val, mid+1, right)
	}
}

func binarySearchLoop(arr []int, val int) []int {
	l, r := 0, len(arr)-1

	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == val {
			return []int{arr[mid]}
		} else if arr[mid] < val {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return []int{}
}

func main() {
	data = []int{10, 3, 5, 6, 34, 6, 3}
	sort.Ints(data)
	fmt.Println(binarySearch(data, 34, 0, len(data)))
	fmt.Println(binarySearchLoop(data, 3))
}

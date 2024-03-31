package main

import (
	"fmt"
	"math/rand"
	"slices"
)

var posts = func(n int) []int {
	arr := make([]int, n)

	for i := range n {
		arr[i] = i + 1
	}
	for i := 0; i < n-1; i++ {
		current := arr[i]
		randomIndex := rand.Int() % n
		randomPos := arr[randomIndex]
		arr[i], arr[randomIndex] = swap(current, randomPos)

	}
	return arr
}
var swap = func(a, b int) (int, int) {
	a ^= b
	b ^= a
	a ^= b
	return a, b
}

var filterPosts = func(arr []int, relevance int) []int {
	result := make([]int, 0, 1)

	for i := range len(arr) {
		if arr[i] > relevance {
			result = append(result, arr[i])
		}
	}
	swapped := make(map[int]bool, len(result))
	slices.Sort(result)
	slices.Reverse(result)

	for i := range len(result) - 1 {
		chance := rand.Float32()
		if chance > .9 && !(swapped[result[i]] && swapped[result[i+1]]) {
			fmt.Println("before swap", result)
			result[i], result[i+1] = swap(result[i], result[i+1])
			fmt.Println("after swap", result)
			swapped[result[i]] = true
			swapped[result[i+1]] = true
		}

	}

	return result
}

func main() {
	arr := posts(10)
	fmt.Printf("filtered posts by relevance %v", filterPosts(arr, 6))

}

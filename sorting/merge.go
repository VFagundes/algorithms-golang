package main

import (
	"fmt"
	"math/rand"
)

const TOTAL_SIZE = 10

var data = func() *[]int {

	L := make([]int, TOTAL_SIZE)
	for i := TOTAL_SIZE - 1; i >= 0; i-- {
		L[i] = rand.Intn(100)
	}
	return &L
}

func mergeSort(items []int) *[]int {
	if len(items) < 2 {
		return &items
	}
	L := mergeSort(items[:len(items)/2])
	R := mergeSort(items[len(items)/2:])
	return merge(*L, *R)
}

func merge(L, R []int) *[]int {
	sorted := make([]int, 0)

	var i, j int = 0, 0
	for i < len(L) && j < len(R) {
		if L[i] < R[j] {
			sorted = append(sorted, L[i])
			i++
		} else {
			sorted = append(sorted, R[j])
			j++
		}
	}
	for ; i < len(L); i++ {
		sorted = append(sorted, L[i])
	}
	for ; j < len(R); j++ {
		sorted = append(sorted, R[j])
	}

	return &sorted
}
func main() {
	L := data()
	fmt.Println(L)
	sorted := mergeSort(*L)

	fmt.Println(sorted)
}

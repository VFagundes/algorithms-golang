package main

import (
	"fmt"
	"math/rand"
)

const TOTAL_SIZE = 10

var data = func() *[]int {

	L := make([]int, TOTAL_SIZE)
	for i := TOTAL_SIZE - 1; i >= 0; i-- {
		L[i] = rand.Intn(10)
	}
	return &[]int{9, 8, 88, 7, 6}
}
var c = 0

func mergeSort(items []int) *[]int {
	if len(items) < 2 {
		return &items
	}
	L := mergeSort(items[:len(items)/2])
	R := mergeSort(items[len(items)/2:])
	fmt.Println(c)
	c++

	return merge(*L, *R)
}

func merge(L, R []int) *[]int {
	sorted := make([]int, 0)
	fmt.Print("L > ")
	fmt.Println(L)
	fmt.Print("R > ")
	fmt.Println(R)
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

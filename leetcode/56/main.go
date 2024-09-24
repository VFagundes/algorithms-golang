package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	r := [][]int{intervals[0]}
	for _, interval := range intervals[1:] {
		//[[1,3],[2,6],[8,10],[15,18]]
		n := r[len(r)-1]
		if n[1] >= interval[0] {
			n[1] = max(n[1], interval[1])
		} else {
			r = append(r, interval)
		}
	}

	return r
}
func main() {

	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 3}, {2, 3}, {8, 15}, {15, 18}}))
}

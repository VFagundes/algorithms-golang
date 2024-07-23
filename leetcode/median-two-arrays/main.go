package main

import (
	"fmt"
	"sort"
)

func findMedian(n1, n2 []int) float64 {

	n1 = append(n1, n2...)
	sort.Ints(n1)
	n := len(n1)
	if (n & 1) == 1 {
		return float64(n1[n>>1])
	}
	return float64(n1[n>>1]+n1[(n>>1)-1]) / 2
}
func findMedianOptimal(n1, n2 []int) float64 {
	m, n := len(n1), len(n2)
	merged := make([]int, m+n)
	count := len(merged)
	i, j := 0, 0
	mid := (m + n) >> 1
	for k := 0; k < count; k++ {
		if i < m && (j >= n || n1[i] < n2[j]) {
			merged[k] = n1[i]
			i++
			continue
		}
		merged[k] = n2[j]
		j++
	}

	if (len(merged) & 1) == 0 {
		return float64(merged[mid-1]+merged[mid]) / 2.0
	}
	return float64(merged[mid])

}

func main() {
	fmt.Println(findMedian([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedian([]int{1, 2}, []int{2, 2, 3}))
	fmt.Printf("optimal findMedian %.2f\n", findMedianOptimal([]int{1, 2}, []int{3, 4}))
	fmt.Printf("optimal findMedian %.2f\n", findMedianOptimal([]int{1, 3}, []int{2}))
	fmt.Printf("optimal findMedian %.2f\n", findMedianOptimal([]int{1, 2, 3}, []int{2, 2}))

}

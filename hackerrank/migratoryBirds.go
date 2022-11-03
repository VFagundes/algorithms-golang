// https://www.hackerrank.com/challenges/migratory-birds/problem
package main

import "fmt"

func migratoryBirds(arr []int32) int32 {
	counterMap := make(map[int32]int32)

	for _, v := range arr {
		counterMap[v]++
	}
	var mostFreqBirdId int32
	var freq int32 = 0
	for k, v := range counterMap {
		if v > freq {
			freq = v
			mostFreqBirdId = k
		}
		if k < mostFreqBirdId && v >= freq {
			mostFreqBirdId = k
		}
	}
	return mostFreqBirdId
}
func main() {
	ar := []int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}
	fmt.Println(migratoryBirds(ar))
}

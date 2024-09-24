package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(makingAnagrams("abc", "abcd"))
	fmt.Println(makingAnagrams("cc", "ccc"))
}

func makingAnagrams(s1 string, s2 string) int32 {
	freq := make([]int, 26)

	// Process the first string: increment the count
	for _, char := range s1 {
		freq[char-'a']++
	}

	for _, char := range s2 {
		freq[char-'a']--
	}

	// Calculate the total number of deletions required
	deletions := 0
	for _, count := range freq {
		deletions += int(math.Abs(float64(count)))
	}

	return int32(deletions)
}

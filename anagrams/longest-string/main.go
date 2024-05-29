package main

import (
	"fmt"
)

// pwwkew
//
// w:
//
//p:0
//w:1 r:1 max:2 (pw)
//w:w r:3
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	l, max := 0, 0
	for r := 0; r < len(s); r++ {
		if i, ok := m[s[r]]; ok && i >= l {
			l = i + 1
		}
		m[s[r]] = r
		if r-l+1 > max {
			max = r - l + 1
		}
	}

	return max
}

func main() {
	fmt.Println("dvdf", lengthOfLongestSubstring("dvdf"))
	fmt.Println("aab", lengthOfLongestSubstring("aab"))
	fmt.Println("abba", lengthOfLongestSubstring("abba"))
	fmt.Println("abcabcbb", lengthOfLongestSubstring("abcabcbb"))
	fmt.Println("pwwkew", lengthOfLongestSubstring("pwwkew"))
	fmt.Println("bbbbb", lengthOfLongestSubstring("bbbbb"))
	fmt.Println("baa", lengthOfLongestSubstring("bbbbb"))
	fmt.Println(" ", lengthOfLongestSubstring("bbbbb"))

}

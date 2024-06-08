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
		if v, ok := m[s[r]]; ok && l <= v {
			l = v + 1
		}
		m[s[r]] = r
		if max < r-l+1 {
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

/*
dvdf 3
aab 2
abba 2
abcabcbb 3
pwwkew 3
bbbbb 1
baa 1
  1
*/

// https://www.youtube.com/watch?v=y2BD4MJqV20
package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("babad"))
}

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}
	l, r := 0, 0
	for i := 0; i < len(s); i++ {
		l1 := expand(s, i, i)
		l2 := expand(s, i, i+1)
		maxLen := l1
		if l1 < l2 {
			maxLen = l2
		}
		if maxLen > r-l {
			l = i - ((maxLen - 1) / 2)
			r = i + (maxLen / 2)
		}

	}
	return s[l : r+1]
}

func expand(s string, l, r int) int {
	if l > r || s == "" {
		return 0
	}

	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}

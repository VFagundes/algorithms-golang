package main

import (
	"fmt"
)

func preProcess(s string) string {
	if len(s) == 0 {
		return "^$"
	}
	res := "^"
	for i := 0; i < len(s); i++ {
		res += "#" + string(s[i])
	}
	res += "#$"
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// findAllPalindromes uses Manacher's Algorithm to find all palindromic substrings.
func findAllPalindromes(input string) []string {
	transformed := preProcess(input)
	fmt.Println(transformed)
	transformedLength := len(transformed)
	palindromeRadii := make([]int, transformedLength)
	centerIndex := 0
	rightBoundary := 0
	var palindromes []string
	// ^#a#a#b#b#b#a#a#$

	for i := 1; i < transformedLength-1; i++ {
		mirrorIndex := 2*centerIndex - i // Index mirrored across the center

		if i < rightBoundary {
			palindromeRadii[i] = min(rightBoundary-i, palindromeRadii[mirrorIndex])
		}

		// Try to expand palindrome centered at i
		for transformed[i+1+palindromeRadii[i]] == transformed[i-1-palindromeRadii[i]] {
			fmt.Println(palindromeRadii)
			palindromeRadii[i]++

		}

		// Adjust center and right boundary if the palindrome extends past the current boundary
		if i+palindromeRadii[i] > rightBoundary {
			centerIndex = i
			rightBoundary = i + palindromeRadii[i]
		}

		// Check if the palindrome has a length greater than 1 and store it
		if palindromeRadii[i] > 1 {
			startIndex := (i - palindromeRadii[i]) / 2 // Convert back to original string index
			palindrome := input[startIndex : startIndex+palindromeRadii[i]]
			palindromes = append(palindromes, palindrome)
		}
	}

	return palindromes
}

// main is the entry point of the program
func main() {
	input := "aabbbaa"
	palindromes := findAllPalindromes(input)
	fmt.Println("Palindromes found:", palindromes)
}

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func reverseString(input string) string {

	runes := []rune(input)

	word := []rune{}
	result := strings.Builder{}

	for i, r := range runes {
		if unicode.IsLetter(r) {
			word = append([]rune{r}, word...)
		} else {
			result.WriteString(string(word))
			result.WriteRune(r)
			word = []rune{}
		}
		if i == len(runes)-1 {
			result.WriteString(string(word))
		}
	}
	return result.String()
}

func main() {

	input := "Hello, Vinicius!|His name was Jean-Francois."
	for _, v := range strings.Split(input, "|") {
		fmt.Println(reverseString(v))
	}

}

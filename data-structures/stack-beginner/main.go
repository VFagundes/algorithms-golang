package main

import (
	"fmt"
	"strconv"
	"strings"
)

var exec = func(strVec []string, fn func(d string) string) {
	for _, v := range strVec {
		fmt.Println(fn(v))
	}
	fmt.Println("---------------------\n\n")
}

func main() {

	exec(reverseMock(), reverseStringStack)
	exec(validParenthesesMock(), validParentheses)
	exec(postFixMock(), postFix)
}

func reverseStringStack(s string) string {
	n := len(s)
	sb := strings.Builder{}
	sb.Grow(n)
	for ; n > 0; n-- {
		sb.WriteByte(s[n-1])
	}
	return sb.String()
}
func validParentheses(s string) string {
	sb := strings.Builder{}
	sb.Grow(len(s))
	sb.WriteString(s)
	sb.WriteString(" is: ")
	st := make([]rune, 0)
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			st = append(st, c)
		} else {
			if len(st) == 0 {
				sb.WriteString("unbalanced\n")
				return sb.String()
			}
			current := st[len(st)-1]
			st = st[:len(st)-1]
			switch current {
			case '(':
				{
					if c != ')' {
						sb.WriteString("unbalanced\n")
						return sb.String()
					}
				}
			case '[':
				{
					if c != ']' {
						sb.WriteString("unbalanced\n")
						return sb.String()
					}
				}
			case '{':
				{
					if c != '}' {
						sb.WriteString("unbalanced\n")
						return sb.String()
					}
				}
			}
		}
	}
	if len(st) > 0 {
		sb.WriteString("unbalanced\n")
	} else {
		sb.WriteString("balanced\n")
	}
	return sb.String()
}
func postFix(s string) string {
	var isNum = func(c byte) bool {
		_, err := strconv.Atoi(string(c))
		return err == nil

	}
	sb := strings.Builder{}
	st := make([]byte, 0)
	for i := 0; i < len(s)-1; i++ {
		// fmt.Print(string(s[i]), " ", isNum(s[i]))
		st = append(st, s[i])
		next := s[i+1]
		if isNum(next) {

		}

	}
	return sb.String()
}

var reverseMock = func() []string {
	return strings.Split("hello,rab,iba", ",")
}
var validParenthesesMock = func() []string {
	return strings.Split("([]),[{]},((())),({}[])(),({[}]),){", ",")
}
var postFixMock = func() []string {
	return strings.Split("3 * 4 + 2,( 1 + 2 ) * ( 3 + 4 ),3 + 4 * 2 / ( 1 - 5 ) ^ 2 ^ 3,5 * 4 - 6 / 2,4 * ( 2 + 3 )", ",")
}

/*
Input: "3 + 4 * 2 / ( 1 - 5 ) ^ 2 ^ 3"
Output: "3 4 2 * 1 5 - 2 3 ^ ^ / +"

Input: "5 * 4 - 6 / 2"
Output: "5 4 * 6 2 / -"

Input: "4 * ( 2 + 3 )"
Output: "4 2 3 + *"

Input: "( 1 + 2 ) * ( 3 + 4 )"
Output: "1 2 + 3 4 + *"

Input: "3 * 4 + 2"
Output: "3 4 * 2 +"
*/

package main

import (
	"fmt"
	"strings"
)

var exec = func(strVec func() []string, fn func(d string) string) {
	for _, v := range strVec() {
		fmt.Println(fn(v))
	}
	fmt.Println("---------------------\n\n")
}

func main() {

	exec(reverseMock, reverseStringStack)
	exec(validParenthesesMock, validParentheses)
	exec(reverseInParenthesesMock, reverseInParentheses)
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
	if len(s) < 2 {
		s := fmt.Sprintf("%s is unbalanced", s)
		return s
	}
	sb := strings.Builder{}
	st := make([]rune, 0)
	sb.WriteString(s)
	sb.WriteString(" is: ")
	for _, c := range s {
		if c == '{' || c == '[' || c == '(' {
			st = append(st, c)
		} else {
			if len(st) == 0 {
				sb.WriteString("unbalanced")
				return sb.String()
			}
			d := st[len(st)-1]
			st = st[:len(st)-1]
			if d == '{' {
				if c != '}' {
					sb.WriteString("unbalanced")
					return sb.String()
				}
			}
			if d == '[' {
				if c != ']' {
					sb.WriteString("unbalanced")
					return sb.String()
				}
			}
			if d == '(' {
				if c != ')' {
					sb.WriteString("unbalanced")
					return sb.String()
				}
			}
		}

	}
	if len(st) > 0 {
		sb.WriteString("unbalanced")
	} else {
		sb.WriteString("balanced")
	}
	return sb.String()
}
func reverseInParentheses(s string) string {
	sb := strings.Builder{}
	return sb.String()
}

var reverseMock = func() []string {
	return strings.Split("hello,rab,iba", ",")
}
var validParenthesesMock = func() []string {
	return strings.Split("([]),[{]},((())),({}[])(),({[}]),){", ",")
}

var reverseInParenthesesMock = func() []string {
	return strings.Split("(bar),foo(bar)baz,foo(bar)baz(blim),foo(bar(baz))blim")

	// For inputString = "(bar)", the output should be
	// solution(inputString) = "rab";
	// For inputString = "foo(bar)baz", the output should be
	// solution(inputString) = "foorabbaz";
	// For inputString = "foo(bar)baz(blim)", the output should be
	// solution(inputString) = "foorabbazmilb";
	// For inputString = "foo(bar(baz))blim", the output should be
	// solution(inputString) = "foobazrabblim".
	// Because "foo(bar(baz))blim" becomes "foo(barzab)blim" and then "foobazrabblim".
}

/*
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func solution(inputString string) string {
	stack := []string{}
	result := ""

	for _, char := range inputString {
		if char == ')' {
			reversed := reverseString(stack[len(stack)-1])
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				result += reversed
			} else {
				stack[len(stack)-1] += reversed
			}
		} else if char == '(' {
			stack = append(stack, "")
		} else {
			if len(stack) == 0 {
				result += string(char)
			} else {
				stack[len(stack)-1] += string(char)
			}
		}
	}

	return result
}
*/

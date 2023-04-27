package main

import (
	"fmt"
	"strings"
)

var exec = func(strVec []string, fn func(d string) string) {
	for _, v := range strVec {
		fmt.Println(fn(v))
	}
}

func main() {

	exec(reverseMock(), reverseStringStack)
	exec(validParenthesesMock(), validParentheses)
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
		if c == '(' {
			st = append(st, c)
		} else if c == '{' {
			st = append(st, c)
		} else if c == '[' {
			st = append(st, c)
		} else {
			if len(st) == 0 {
				sb.WriteString("unbalanced\n")
				return sb.String()
			}
			current := st[len(st)-1]
			st = st[:len(st)-1]
			if current == '(' {
				if c != ')' {
					sb.WriteString("unbalanced\n")
					return sb.String()
				}
			}
			if current == '[' {
				if c != ']' {
					sb.WriteString("unbalanced\n")
					return sb.String()
				}
			}
			if current == '{' {
				if c != '}' {
					sb.WriteString("unbalanced\n")
					return sb.String()
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

var reverseMock = func() []string {
	return strings.Split("hello,rab,iba", ",")
}
var validParenthesesMock = func() []string {
	return strings.Split("([]),[{]},((())),({}[])(),({[}]),){", ",")
}

/*
Input: "({}[])()"

Expected Output: true

Input: "({[}])"

Expected Output: false

Input: "(a + b) * (c - d)"

Expected Output: true

Input: "{()[]"

Expected Output: false

Input: "(())({})"

Expected Output: true

Input: "({)}"
*/

/*
Reverse a string using a stack

Input: "hello"

Output: "olleh"

Check if a string of parentheses is balanced

Input: "({}[])()"

Output: true

Input: "({[}])"

Output: false

Evaluate a postfix expression using a stack

Input: "3 4 + 5 *"

Output: 35

Input: "4 2 * 5 + 1 -"

Output: 8

Implement the Tower of Hanoi problem using stacks

Input: Three stacks representing the pegs with n disks on the first peg.

Output: The final configuration of disks on the third peg.
*/

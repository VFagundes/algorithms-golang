package main

import (
	"fmt"
	"strings"
)

var d = func() []string {
	// Input: str = “(skeeg(for)skeeg)”
	// Output: geeksforgeeks

	// Input: str = “((ng)ipm(ca))”
	// Output: camping

	// Input: (bar)
	// Output: rab

	return strings.Split("(bar),(skeeg(for)skeeg),((ng)ipm(ca))", ",")
}

func main() {
	fmt.Println("akaa")
}

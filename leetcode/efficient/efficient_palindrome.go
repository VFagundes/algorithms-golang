package main

import (
	"reflect"
)

func findAllPalindromes(input string) []string {
	result := make([]string, 1)


	return result
}


func main() {
	testInput := "abbaeae"
	expectedOutput := []string{"abba", "bb", "aea", "eae"}

	result := findAllPalindromes(testInput)
	if !reflect.DeepEqual(result, expectedOutput) {
		panic("Test failed: Expected output does not match actual output")
	}
}

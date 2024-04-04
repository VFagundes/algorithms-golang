// https://www.hackerrank.com/challenges/cats-and-a-mouse/
package main

import "fmt"

func catAndMouse(x int32, y int32, z int32) string {
	// A custom absolute difference function that avoids float conversion.
	absDiff := func(a, b int32) int32 {
		if a > b {
			return a - b
		}
		return b - a
	}

	xToZ := absDiff(z, x)
	yToZ := absDiff(z, y)

	if xToZ == yToZ {
		return "Mouse C"
	} else if yToZ < xToZ {
		return "Cat B"
	}
	return "Cat A"
}

func main() {

	fmt.Println(catAndMouse(10, 20, 15))
}

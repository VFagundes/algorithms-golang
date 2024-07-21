package main

import "fmt"

func main() {

	d := make([]int, 5)
	d[len(d)-1] = 33
	d[0] = 1
	fmt.Println(d[:3])
	fmt.Println(d[0:4])
	fmt.Println(d[:])
	fmt.Println(d[3 : len(d)-1])

}

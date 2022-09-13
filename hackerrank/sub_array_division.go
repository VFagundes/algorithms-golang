package main

import (
	"fmt"
	"strconv"
	"strings"
)

// sub array division https://www.hackerrank.com/challenges/the-birthday-bar/problem?isFullScreen=true

func prepareArray(d string) []int {
	data := strings.Split(d, " ")
	s := make([]int, len(data))
	for i, v := range data {
		num, _ := strconv.Atoi(v)
		s[i] = int(num)
	}

	return s
}
func exec(ds ...string) {
	for _, v := range ds {
		line := strings.Split(v, "\n")
		s := prepareArray(line[0])
		d, m := prepareLimits(line[1])
		fmt.Println(birthday(s, d, m))
	}
}
func prepareLimits(d string) (int, int) {
	s := strings.Split(d, " ")
	d1, _ := strconv.Atoi(s[0])
	d2, _ := strconv.Atoi(s[1])
	return d1, d2
}

// let cont = 0

// s.forEach((num, indice) => {
// 	let sum = 0
// 	for(let i = indice; i < indice + m; i++){
// 		sum += s[i]
// 	}
// 	if(sum === d) cont++
// })
// return cont

func birthday(s []int, d int, m int) int {
	total := 0

	for i := 0; i <= len(s)-m; i++ {
		sum := 0
		for j := i; j < i+m; j++ {
			sum += s[j]
		}
		if sum == d {
			total++
		}
	}
	return total

}

func main() {

	d1 := `1 2 1 3 2
3 2`
	d2 := `1 1 1 1 1 1
3 2`
	d3 := `4
4 1`
	d4 := `2 5 1 3 4 4 3 5 1 1 2 1 4 1 3 3 4 2 1
18 7`
	fmt.Println(d1, d2, d3, d4)
	exec(d1, d2, d3, d4)
	// exec(d4)
}

package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		var fizz bool = i%3 == 0
		var buzz bool = i%5 == 0
		fmt.Println(i)
		if fizz && buzz {
			fmt.Println("FizzBuzz")
		} else if fizz {
			fmt.Println("Fizz")
		} else if buzz {
			fmt.Println("Buzz")
		}
	}

}

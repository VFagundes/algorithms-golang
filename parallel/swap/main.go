package main

import (
	"fmt"
	"sync"
)

func Swap(a, b int) (int, int) {
	return b, a
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(a, b int) {
			defer wg.Done()
			j, k := Swap(i, i+1)
			fmt.Println(j, k)
		}(i, i+1)
	}
	wg.Wait()
}

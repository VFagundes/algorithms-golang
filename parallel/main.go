package main

import (
	"fmt"
	"time"
)

func receiveData(ch chan<- string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		ch <- fmt.Sprintf("Data %d", i)
	}
	close(ch)
}

func printData(ch <-chan string) {
	for data := range ch {
		fmt.Println(data)
	}
}

func main() {
	dataChannel := make(chan string)

	go receiveData(dataChannel)
	go printData(dataChannel)

	// Wait for goroutines to finish
	time.Sleep(6 * time.Second)
}

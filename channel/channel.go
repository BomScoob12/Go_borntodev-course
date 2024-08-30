package main

import (
	"fmt"
	"time"
)

func processChannel(c chan string, data string) {
	// send data to channel
	c <- data
}

func main() {
	// create a channel
	demoChannel := make(chan string)
	// send data to channel
	// what is go command?
	// go command is used to run a goroutine. 
	// Goroutine is a function or method which executes independently 
	// and simultaneously in connection with any other goroutines 
	// present in your program. Or in other words, 
	// every concurrently executing activity in Go language is known as a goroutine.
	go processChannel(demoChannel, "Lorem")
	// receive data from channel
	fmt.Println("Data 1 sent to channel")
	fmt.Println("Data 1 received from channel: ", <-demoChannel)
	time.Sleep(2 * time.Second)
}
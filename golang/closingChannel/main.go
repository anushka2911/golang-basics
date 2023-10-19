package main

import "fmt"

func main() {
	fmt.Printf("Beginning of the program\n")

	channel := make(chan string, 1)
	channel <- "Hello"
	close(channel)
	temp := <-channel
	fmt.Printf("The value received from the channel is %s\n", temp)
	// channel <- "World" // this gives us panic
	// temp = <-channel
	fmt.Printf("The value received from the channel is %s\n", temp)
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hello, World!")

	//channel := make(chan int, 2) // bidirectional channel

	// go func() { //this is an anonymous function
	// 	channel <- 42 // send
	// }() // () this is to call the anonymous function

	//temp := <-channel // receive

	//fmt.Println(temp)
	//var wg sync.WaitGroup
	//wg.Add(1)
	// Channel := make(chan int)

	// go func() {
	// 	fmt.Printf("Checking for send only channel /n")
	// 	Channel <- 42
	// }()

	// <-Channel // this is a blocking call
	// //and it will remain blocked until a value is sent to the channel

	// var wg sync.WaitGroup
	// wg.Add(1)
	// send_only_channel := make(chan<- int, 1) //buffered channel
	// go func() {
	// 	defer wg.Done()
	// 	send_only_channel <- 104
	// }()
	// wg.Wait()

	var wg sync.WaitGroup
	wg.Add(1)
	receive_only_channel := make(<-chan int, 1) //buffered channel
	go func() {
		defer wg.Done()
		fmt.Println(<-receive_only_channel)
	}()
	wg.Wait()
}

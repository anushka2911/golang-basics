package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Printf("Beginning of the main function\n")
	var wg sync.WaitGroup
	wg.Add(1)

	channel := make(chan int)
	go addTwoNumbers(10, 20, channel, &wg)
	result := <-channel // this is a blocking call
	close(channel)
	wg.Wait()

	fmt.Println("Result is ", result)
	fmt.Println("End of the main function")
}

func addTwoNumbers(a int, b int, channel chan int, wg *sync.WaitGroup) {
	defer wg.Done() // this specifies that the goroutine is completed
	fmt.Printf("Beginning of the addTwoNumbers function\n")
	sum := a + b
	channel <- sum
	fmt.Println("End of the addTwoNumbers function")
}

// The main goroutine launches the addTwoNumbers goroutine using go addTwoNumbers(10, 20, channel, &wg).

// The addTwoNumbers goroutine executes concurrently with the main goroutine.

// Inside the addTwoNumbers goroutine, it calculates the sum, sends it to the channel (channel <- sum), and then proceeds to print messages.

// Meanwhile, the main goroutine is blocked at the line result := <-channel until a value is received from the channel.

// Once the addTwoNumbers goroutine sends the sum to the channel, the main goroutine proceeds to execute the line result := <-channel, receiving the value from the channel.

// After receiving the value, the main goroutine can continue with other statements, such as closing the channel (close(channel)) and waiting for the addTwoNumbers goroutine to complete (wg.Wait()).

// So, yes, the result := <-channel line is executed as soon as a value is available in the channel, allowing for concurrent execution with the addTwoNumbers goroutine.

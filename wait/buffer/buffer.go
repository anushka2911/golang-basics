package main

import "fmt"

func main() {
	fmt.Println("Beginning of the main function")
	fmt.Println("Exp with gcommands")
	channel := make(chan int, 2)

	channel <- 10
	channel <- 20
	//channel <- 30 //this line will cause a deadlock as the channel capacity is 2

	for i := range channel {
		fmt.Println(i)
	}
}

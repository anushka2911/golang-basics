package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Printf("Beginning of the main function\n")
	var wg sync.WaitGroup
	wg.Add(1) //this specifies that it has 1 goroutine to wait for
	go greeting("Good Morning Anushka", &wg)
	wg.Wait() // waits for the go routine to complete
	//the below statement is executed after 5seconds

	fmt.Printf("End of the main function\n")

}

func greeting(s string, wg *sync.WaitGroup) {
	fmt.Printf("This is a greeting message %s \n", s)
	wg.Done() // this specifies that the goroutine is completed
}

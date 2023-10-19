package main

import "fmt"

func printGreetings(temp ...string) {
	for _, val := range temp {
		println(val)
	}
	fmt.Println("\n")
}
func main() {
	printGreetings("Hello")
	printGreetings("Hello", "World")
	printGreetings("Hello", "World", "How", "are", "you")
}

package main

import "fmt"

type person struct {
	name string
	age  int
}

// This is a method associated with the person struct
func (p person) DoubleAge() int {
	return p.age * 2
}

func main() {
	// struct in golang
	// custom data type, different types of data types are combined under one field like a class
	p := person{ // assigning values to struct
		age:  20,
		name: "Anushka",
	}
	fmt.Println(p)

	// methods in golang
	// Methods are functions that are associated with a struct
	doubleAge := p.DoubleAge()
	fmt.Println(doubleAge)

	//difference between method and function
	// function is a block of code that can be called by its name
	// method is a function that is associated with a struct so after writing func we have to write the parent struct name and then the method name

	//is golang fast?
	// yes, golang is fast because it is a compiled language and it is statically typed and it has a garbage collector and it has a very good concurrency model and it has a very good standard library

	//slice in golang
	//slice is like an array but with dynamic size
	//var anushka []int

}

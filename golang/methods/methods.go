package main

import "fmt"

type Person struct {
	Name      string
	Age       int
	HairColor string
	Batch     int
	Mobile    int
}

func (p Person) PrintName() {
	fmt.Println("inside print name")
	fmt.Println(p.Name)
}

func main() {

	Anushka := Person{
		Name:      "anushee",
		Age:       21,
		HairColor: "black",
		Batch:     2019,
		Mobile:    1234,
	}
	Anushka.PrintName()
}

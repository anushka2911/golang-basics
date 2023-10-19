package main

import "fmt"

type Person struct {
	Name      string
	Age       int
	HairColor string
	Batch     int
}

type AboutPerson interface {
	UserDetails()
}

func (p Person) UserDetails() {
	fmt.Println("inside user details")
	fmt.Println(p.Name)
	fmt.Println(p.Age)
	fmt.Println(p.HairColor)
}

func main() {
	anushka := Person{
		Name:      "anushee",
		Age:       21,
		HairColor: "black",
		Batch:     2019,
	}

	anushka.UserDetails()
}

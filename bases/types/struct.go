package main

import (
	"fmt"
)

type product struct {
	name        string
	description string
	price       float64
}

type user struct {
	firstname string
	lastname  string
	age       uint8
}

func (p product) display() string {
	return p.name + p.description
}

func (u user) sayHello() string {
	return "Hello" + u.firstname
}

func main() {
	u := user{
		firstname: "John",
		lastname:  "Doe",
		age:       30,
	}
	p := product{
		name:        "xbox",
		description: "one x",
		price:       500.99,
	}
	fmt.Println(p.display())
	fmt.Println(u.sayHello())
	fmt.Println(u.firstname, u.lastname, u.age)

}

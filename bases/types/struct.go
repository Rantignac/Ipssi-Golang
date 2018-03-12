package main

import (
	"fmt"
)

type user struct {
	firstname string
	lastname  string
	age       uint8
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
	fmt.Println(u.sayHello())
	fmt.Println(u.firstname, u.lastname, u.age)

}

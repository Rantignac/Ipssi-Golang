package main

import (
	"fmt"
)

func GetNames() (string, string) {
	return "John", "Doe"
}

func ShowNames(firstname, lastname string) {
	fmt.Println(firstname, lastname)

}

type Funky func()

func (f Funky) Dance() {
	f()
}

type Dancer interface {
	Dance()
}
type Str string

func (s Str) Dance() {
	println(s)
}

func StartParty(d Dancer) {
	d.Dance()

}

func main() {

	ShowNames(GetNames())
	var s Str
	s = "je suis une chaine de caractère et je dance.."
	var f Funky
	f = func() {
		fmt.Println("I'm dancing Get Funky")
	}

	StartParty(f)
	StartParty(s)
}

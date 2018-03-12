package main

import (
	"fmt"
)

const language = "Go"

//a == 0
//b == 1
//c == 2

const (
	a = 2 * iota
	b
	c
)

func main() {

	fmt.Println(language, a, b, c)

}

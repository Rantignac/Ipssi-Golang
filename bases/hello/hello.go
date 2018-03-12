package main

import (
	"fmt"
	"os"
)

func main() {
	var name string = "World"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	fmt.Printf("Hello, %s", name)

}

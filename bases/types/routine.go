package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("A Long Task")
	}()
	println("Done.")
	time.Sleep(time.Second * 3)
}

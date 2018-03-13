package main

import (
	"fmt"
)

func main() {
	m := map[string]string{
		"name":    "John",
		"country": "UK",
	}

	if v, ok := m["name"]; ok {
		fmt.Println(v, m["country"])

	}

	m2 := map[string]interface{}{
		"prenom": "John",
		"nom":    "Smith",
		"age":    35,
		"size":   2.66,
	}

	fmt.Println(m2)

}

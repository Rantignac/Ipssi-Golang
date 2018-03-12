package main

import (
	"fmt"
)

var (
	ecole   = "ipssi"
	adresse = "252 rue Claude Tilier"
)

func main() {
	var name string
	name = "Gopher"

	var x, y int
	x = 10
	fmt.Println(name, x, y)

	//Déclaration de variable avec un type dynamique
	a := 60
	a = 20

	w, z := 5, false

	//Le Swap
	u, v := 10, 35
	v, u = u, v

	fmt.Println(a, w, z)

	version := string("ma valeur")

	fmt.Println(version)
}

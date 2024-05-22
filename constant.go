package main

import (
	"fmt"
)

func main() {
	const firstname string = "M. Aji"
	const lastname = "Perdana"

	// error, because constant cannot be reassigned
	// firstname = "Frasiska"
	// fmt.Println(firstname)

	const (
		name string = "M. Aji Perdana"
		age int = 20
	)

	fmt.Println(name)
	fmt.Println(age)
}
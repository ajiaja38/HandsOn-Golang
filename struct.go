package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var aji Customer

	aji.Name = "M. Aji Perdana"
	aji.Address = "Lampung"
	aji.Age = 22

	frasiska := Customer{
		Name:    "Frasiska Risma Yolanda",
		Address: "Wates Selatan, Pringsewu",
		Age:     22,
	}

	alam := Customer{"Ridwan Alamsyah", "Tangerang", 8}

	fmt.Println(aji)
	fmt.Println(frasiska)
	fmt.Println(alam)
}

package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func sayHellos(customer Customer) string {
	return "Hello " + customer.Name
}

func main() {
	vera := Customer{
		Name:    "Vera",
		Address: "Jakarta",
		Age:     20,
	}

	Aji := Customer{"M. Aji Perdana", "Lampung", 22}

	fmt.Println(vera)
	fmt.Println(Aji)
	fmt.Println(sayHellos(vera))
}

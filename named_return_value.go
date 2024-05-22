package main

import "fmt"

func get_complete_name() (firstName string, midleName string, lastName string) {
	firstName = "Frasiska"
	midleName = "Risma"
	lastName = "Yolanda"

	return firstName, midleName, lastName
}

func main() {
	a, b, c := get_complete_name()

	fmt.Println(a, b, c)
}

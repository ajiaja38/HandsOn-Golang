package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "M. Aji Perdana",
		"address": "Jakarta",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])

	book := make(map[string]string)
	book["title"] = "buku golang"
	book["author"] = "M. Aji Perdana"
	book["wrong"] = "oops"

	fmt.Println(book)

	delete(book, "wrong")

	fmt.Println(book)
}

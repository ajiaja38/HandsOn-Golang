package main

import "fmt"

func getFullName() (string, string) {
	return "M. Aji", "Perdana"
}

func main() {
	firstName, lastName := getFullName()

	result := fmt.Sprintf("Hallo My Name is %s %s", firstName, lastName)
	fmt.Println(result)

	// menghiraukan return value
	first, _ := getFullName()
	fmt.Println(first)
}

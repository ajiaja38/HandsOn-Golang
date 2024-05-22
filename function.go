package main

import "fmt"

func sayHello(firstName string, lastName string) string {
	return fmt.Sprintf("Hello %s %s", firstName, lastName)
}

func main() {
	result := sayHello("M. Aji", "Perdana")
	fmt.Println(result)
}

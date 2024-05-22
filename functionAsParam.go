package main

import "fmt"

func filteredName(name string, filter func(name string) string) {
	filtered := filter(name)

	fmt.Println("Hello: ", filtered)
}

func spamFilter(name string) string {
	if name == "Aji" {
		return name
	} else {
		return "Salah Orang!"
	}
}

func main() {
	filter := spamFilter

	filteredName("siska", filter)
}

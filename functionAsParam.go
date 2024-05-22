package main

import "fmt"

func filteredName(name string, filter func(name string) string) {
	filtered := filter(name)

	fmt.Println(filtered)
}

func spamFilter(name string) string {
	if name == "Aji" {
		return name
	} else {
		return "Wrong Person"
	}
}

func main() {
	filter := spamFilter

	filteredName("Aji", filter)
}

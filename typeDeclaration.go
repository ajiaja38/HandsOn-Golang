package main

import "fmt"

type Filter func(name string) string

func filteredValue(name string, filter Filter) string {
	result := filter(name)

	return result
}

func spamFilterValue(name string) string {
	if name == "Anjing" || name == "Kadal" {
		return "..."
	} else {
		return name
	}
}

func main() {
	filter := spamFilterValue

	result := filteredValue("Anjing", filter)
	fmt.Println(result)
}

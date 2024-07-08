package main

import "fmt"

type Blacklist func(name string) bool

func isBlacklist(name string, blacklist Blacklist) string {
	if blacklist(name) {
		return "kata kata kamu telah terblacklist"
	} else {
		return "kata kata anda lolos"
	}
}

func main() {
	blackListWord := func(name string) bool {
		return name == "Anjing"
	}

	result := isBlacklist("Anjing", blackListWord)
	results := isBlacklist("Kadal", func(name string) bool {
		return name == "Anjing"
	})

	fmt.Println(result)
	fmt.Println(results)
}

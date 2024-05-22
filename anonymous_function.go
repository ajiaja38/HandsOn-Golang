package main

import "fmt"

type Blacklist func(name string) bool

func isBlacklist(name string, blacklist Blacklist) string {
	if blacklist(name) {
		return "Maaf, Kata-kata anda diblacklist"
	} else {
		return "Selamat, kata kata anda lulus sensor"
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "Anjing" || name == "Kadal"
	}

	result := isBlacklist("Anjing", blacklist)

	results := isBlacklist("Kadal", func(name string) bool {
		return name == "Anjing" || name == "Kadal"
	})

	fmt.Println(result)
	fmt.Println(results)
}

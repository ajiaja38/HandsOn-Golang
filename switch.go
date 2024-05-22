package main

import "fmt"

func main() {
	name := "Aji"

	switch name {
	case "Aji":
		fmt.Println("Hello ", name)
		break
	case "Siska":
		fmt.Println("Hello ", name)
		break
	default:
		fmt.Println("Boleh Kenalan?")
	}

	// switch short expressiong
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Terlalu Panjang")
		break
	case false:
		fmt.Println("Benar")
		break
	}
}

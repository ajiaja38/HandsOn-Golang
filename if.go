package main

import "fmt"

func main() {
	name := "Siska"

	if name == "Aji" {
		fmt.Println("Hello,", name)
	} else if name == "Siska" {
		fmt.Println("Hello, Siska")
	} else {
		fmt.Println("Boleh Kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}

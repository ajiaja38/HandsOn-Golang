package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("Masukkan Nama Anda: ")
	fmt.Scanln(&name)

	fmt.Print("Masukkan Umur anda: ")
	fmt.Scanln(&age)

	fmt.Printf("Hello, %s! kamu berumur %d \n", name, age)
}

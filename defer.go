package main

import "fmt"

func logging() {
	fmt.Println("Ini Logging")
}

func getUserById() {
	defer logging()

	fmt.Println("Berhasil get User by id")
}

func main() {
	getUserById()
}

package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	// lebih sederhana lagi
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	for i := 5; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}

		fmt.Println("")
	}

	names := []string{"Aji", "Siska", "Elin", "Vera"}

	fmt.Println()

	// for biasa
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	fmt.Println()

	// dengan range
	for i, name := range names {
		fmt.Println("Index ", i, "=", name)
	}

	fmt.Println()

	// jika tidak butuh indexnya:
	for _, name := range names {
		fmt.Println("Nama: ", name)
	}

}

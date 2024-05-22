package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "M. Aji"
	names[1] = "Perdana"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])

	var values = [3]int{
		90,
		10,
	}

	fmt.Println(values)

	var numberOfAges = [...]int{
		90,
		10,
		20,
		50,
	}

	fmt.Println(numberOfAges)
	fmt.Println(len(numberOfAges))
}

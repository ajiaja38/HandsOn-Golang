package main

import "fmt"

func multipleAll(args ...int) int {
	result := 1

	for _, arg := range args {
		result *= arg
	}

	return result
}

func main() {
	fmt.Println(multipleAll(1, 2, 3, 4))

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(multipleAll(numbers...))
}

package main

import "fmt"

func sumAll(numbers ...int) int {
	result := 0

	for _, number := range numbers {
		result += number
	}

	return result
}

func main() {
	fmt.Println(sumAll(10, 10, 10, 10, 5))
	fmt.Println(sumAll(15, 15, 15, 5))

	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(sumAll(numbers...))
}

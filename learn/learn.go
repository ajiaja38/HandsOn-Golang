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
	var numbers []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Hasil dari pejumlahan adalah = ")
	fmt.Println(sumAll(numbers...))
}

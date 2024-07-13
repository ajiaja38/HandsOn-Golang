package main

import "fmt"

func sumAll(args ...int) int {
	result := 0

	for _, arg := range args {
		result += arg
	}

	return result
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	result := sumAll(numbers...)

	fmt.Println(result)
	fmt.Println(sumAll(1, 2, 3, 4, 5, 6))
}

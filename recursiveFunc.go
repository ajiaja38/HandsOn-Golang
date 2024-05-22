package main

import "fmt"

func factorial(number int) int {
	if number == 1 {
		return 1
	} else {
		return number * factorial(number-1)
	}
}

func main() {
	result := factorial(10)
	fmt.Println(result)
}

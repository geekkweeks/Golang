package main

import "fmt"

func getAreaCalculation(l int, w int) int {
	return l * w
}

func main() {
	area := getAreaCalculation

	result := area(3, 5)
	fmt.Println(result)
}

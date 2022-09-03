package main

import "fmt"

func sumAll(inputs ...int) int {
	total := 0
	for _, value := range inputs {
		total += value
	}
	return total
}

func main() {
	total := sumAll(2, 3, 5, 1)
	fmt.Println(total)

	//send slice param into vararg
	sliceSample := []int{10, 2, 3, 1}
	reTotal := sumAll(sliceSample...)
	fmt.Println(reTotal)
}

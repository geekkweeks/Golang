package main

import (
	"fmt"
	"strconv"
)

func calculate(x int, y int) int {
	return x * y
}

func showResult(input int) string {
	return "The result is " + strconv.Itoa(input)
}

func main() {
	x := 5
	y := 10
	res := calculate(x, y)
	fmt.Println(showResult(res))
}

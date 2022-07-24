package main

import "fmt"

func calculate(x int, y int) {
	fmt.Println(x * y)
}

func main() {
	x := 5
	y := 10
	calculate(x, y)
}

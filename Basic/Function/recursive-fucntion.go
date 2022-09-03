package main

import "fmt"

func recursiveLoop(value int) int {
	if value == 0 { //condition to stop recursive
		return 1
	} else {
		return value * recursiveLoop(value-1)
	}
}

func main() {
	total := recursiveLoop(5)
	fmt.Println(total)
}

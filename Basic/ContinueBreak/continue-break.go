package main

import "fmt"

func main() {

	for x := 0; x <= 5; x++ {

		if x == 3 {
			break
		}
		fmt.Println("Loop - ", x)
	}

	for y := 0; y <= 20; y++ {
		if y%2 == 0 {
			continue
		}
		fmt.Println("val y: ", y)
	}

}

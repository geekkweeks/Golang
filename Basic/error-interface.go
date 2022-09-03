package main

import (
	"errors"
	"fmt"
)

func Div(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Y has to greater than 0")
	} else {
		result := x / y
		return result, nil
	}
}

func main() {
	// var errorSample = errors.New("This is an error")
	// fmt.Println(errorSample.Error())

	result, err := Div(100, 0)
	if err == nil {
		fmt.Println("Result: ", result)
	} else {
		fmt.Println("Error ", err.Error())
	}
}

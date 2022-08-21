package main

import "fmt"

func CustomError(i int) interface{} {
	if i == 1 {
		return i
	} else if i == 2 {
		return i
	} else {
		return "Something went wrong"
	}
}

func main() {
	var data interface{} = CustomError(4)
	fmt.Println(data)
}

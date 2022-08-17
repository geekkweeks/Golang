package main

import "fmt"

func main() {

	name := "alfan"
	modifyName := func() {
		name = "zahri"

	}

	modifyName()
	fmt.Println(name)

}

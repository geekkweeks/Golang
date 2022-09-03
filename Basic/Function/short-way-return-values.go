package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {
	firstName = "Alfan"
	middleName = "zahri"
	lastName = "yono"

	return
}

func main() {
	x, y, z := getFullName()
	fmt.Println(x, y, z)

}

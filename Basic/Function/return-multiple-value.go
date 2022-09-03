package main

import (
	"fmt"
)

func detailInfo(name string, salary int, isPermanent bool) (string, int, bool) {
	return name, salary, isPermanent
}

func fullName() (string, string, string) {
	return "Alfan", "Zahri", "Yono"
}

func main() {
	name, salary, isPermanent := detailInfo("Alfan", 500000, true)
	fmt.Println(name)
	fmt.Println(salary)
	fmt.Println(isPermanent)

	firstName, _, _ := fullName()
	fmt.Println(firstName)

	_, middleName, _ := fullName()
	fmt.Println(middleName)
}

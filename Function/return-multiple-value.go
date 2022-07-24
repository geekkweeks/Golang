package main

import (
	"fmt"
)

func detailInfo(name string, salary int, isPermanent bool) (string, int, bool) {
	return name, salary, isPermanent
}

func main() {
	name, salary, isPermanent := detailInfo("Alfan", 500000, true)
	fmt.Println(name)
	fmt.Println(salary)
	fmt.Println(isPermanent)
}

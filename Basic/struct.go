package main

import "fmt"

/*
- Struct is data template or data prototype
- Define  some fields by using struct
- Can't be used directly
*/
type Employee struct {
	Name, Position string
	Salary         int
	IsFulltime     bool
}

func main() {
	var emp Employee
	emp.Name = "alfan"
	emp.Position = "GO Developer"
	emp.Salary = 3444
	emp.IsFulltime = true

	fmt.Println(emp.Name)
	fmt.Println(emp.Position)
	fmt.Println(emp.Salary)
	fmt.Println(emp.IsFulltime)

	// ===========Struct Literal==========
	empContract := Employee{
		Name:       "Anwar",
		Position:   "FE Developer",
		Salary:     456,
		IsFulltime: false,
	}
	fmt.Println(empContract)

	empFullTime := Employee{"Joko", "BE DEVELOPER", 34343, true}
	fmt.Println(empFullTime)
	// ===================================
}

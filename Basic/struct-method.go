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

func (employee Employee) showEmployeeSalary(bossName string) {
	fmt.Println("Your boss give you salary $", employee.Salary)
}

func (employee Employee) showEmployeeName() {
	fmt.Println("Your name is ", employee.Name)
}

func main() {
	var emp Employee
	emp.Name = "alfan"
	emp.Position = "GO Developer"
	emp.Salary = 3444
	emp.IsFulltime = true

	emp.showEmployeeName()
	emp.showEmployeeSalary("Brian")
}

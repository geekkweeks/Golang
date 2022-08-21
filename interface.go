package main

import "fmt"

// Abstract
type PersonName interface {
	// Contract
	GetName() string
}

func printName(personName PersonName) {
	fmt.Println("HI ", personName.GetName())
}

type Person struct {
	Name string
}

type Car struct {
	Brand string
}

/*
	- to implement interface
		*) function name,param and return should be same with interface(contract)
*/
func (person Person) GetName() string {
	return person.Name
}

/*
	- to implement interface
		*) function name,param and return should be same with interface(contract)
		*) can be implemented with many structs, as long as the function name,param and return same with contract of interface
*/
func (car Car) GetName() string {
	return car.Brand
}

func main() {
	var personData Person
	personData.Name = "Alfan"
	printName(personData)

	carData := Car{
		Brand: "Hyundai IONIC",
	}
	printName(carData)
}

package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	name := "Alfan"
	result := HelloWorld(name)
	if result != "Hello "+name {
		panic("Result is not same")
	}
}

func TestHelloWorldEmployee(t *testing.T) {
	name := "Yono"
	result := HelloWorld(name)
	if result != "Hello "+name {
		panic("Result is not same")
	}
}

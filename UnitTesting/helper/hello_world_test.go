package helper

import (
	"fmt"
	"testing"
)

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

//failed testing. Not recommened to using Panic function to define failed inside of unit testing
func TestHelloWorldFailed(t *testing.T) {
	name := "Yono"
	result := HelloWorld(name)
	if result != "HI "+name {
		t.Fail()
	}
	fmt.Println("Keep executed even the unit test was failed by using t.Fail()")
}

func TestHelloWorldFailedNo(t *testing.T) {
	name := "Yono"
	result := HelloWorld(name)
	if result != "HI "+name {
		t.FailNow()
	}
	fmt.Println("Will not executed")
}

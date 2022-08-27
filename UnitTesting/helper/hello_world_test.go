package helper

import (
	"fmt"
	"testing"
)

// Unit test before and After
// function name should be 'TestMain'
// TestMain can running in one package only
func TestMain(m *testing.M) {
	fmt.Println("Before all unit test running")

	m.Run()

	fmt.Println("After all unit test running")
}

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

// Error and Fatal functions have argument that we can set detail of error.
func TestHelloWorldError(t *testing.T) {
	name := "Yono"
	result := HelloWorld(name)
	if result != "Hi "+name {
		t.Error("The result must be 'Hi " + name)
	}
	fmt.Println("Will executed")
}

func TestHelloWorldFatal(t *testing.T) {
	name := "Yono"
	result := HelloWorld(name)
	if result != "Hi "+name {
		t.Fatal("The result must be 'Hi " + name)
	}
	fmt.Println("Not executed")
}

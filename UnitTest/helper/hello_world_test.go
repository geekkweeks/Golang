package helper

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M){
	fmt.Println("Start testing")

	m.Run()

	fmt.Println("Finish Testing")

}

func TestHelloWorld(t *testing.T){
	result := HelloWorld("zahri")

	expectedResult := "hello Alfan"

	if result != expectedResult  {
		// error(test failed)
		// if line code is error, the line code below is still executed(not  stop the process)
		t.Fail()		
	}

	fmt.Println("Finished")
}


func TestHelloWorld2(t *testing.T){
	result := HelloWorld("zahri")

	expectedResult := "hello Alfan"

	if result != expectedResult  {
		// error(test failed)
		// will stop the process
		t.FailNow()		
	}	
}

/*
// if we want to return error message we can use t.Error(argument)	
// it will not stop the process below even current line is error
// this is recommended than Fail function. because we can return error message
*/
func TestHelloWorldError(t *testing.T){
	result := HelloWorld("zahri")

	expectedResult := "hello Alfan"

	if result != expectedResult  {
		// error(test failed)
		t.Error("expected result is " + expectedResult)		
	}	
	fmt.Println("TestHelloWorldError is done")
}


/*
// if we want to return error message we can use t.Error(argument)	
// it will stop the process below even current line is error
// this is recommended than Fail function. because we can return error message
*/
func TestHelloWorldFatal(t *testing.T){
	result := HelloWorld("zahri")

	expectedResult := "hello Alfan"

	if result != expectedResult  {
		// error(test failed)
		// will stop the process
		t.Fatal("expected result is " + expectedResult)		
	}	
	fmt.Println("TestHelloWorldFatal is done")
}
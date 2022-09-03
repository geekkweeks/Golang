package main

import "fmt"

func logging(applicationName string) {
	fmt.Println("Application: ", applicationName, " has completed")
	errorMessage := recover() //will catch an error from panic
	fmt.Println("Application has an error message: ", errorMessage)
}

func executeTask(isError bool) {
	defer logging("Defer App")
	if isError {
		panic("Application was error") // if panic is called, the code below will not be executed
	}
	fmt.Println("Finished")
}

func main() {
	executeTask(true)
	fmt.Println("The application will not stop even though has an error")
}

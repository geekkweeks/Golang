package main

import "fmt"

/*
	- Convert from empty interface into what we want
	- make sure that what we want to convert from empty interface into our goal is same
	- for example empty interface return string, so we have to convert into string too. if not, the program will return panic/error
	- To avoid panic in type assertion, we can use Type Assertion Switch. So the application will not crash
*/

func wentWrong() interface{} {
	return false
}

func main() {
	var res interface{} = wentWrong()
	/*
		- funct wenWrong return empty interface with string
		- if we convert beside string, the program will return an error/ Panic
	*/
	//convert empty interfae into string
	// var resString = res.(string)
	// fmt.Println(resString)

	switch value := res.(type) {
	case string:
		fmt.Println("value", value, " is string")
	case int:
		fmt.Println("value", value, " is int")
	default:
		fmt.Println("Unknown type")

	}
}

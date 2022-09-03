package main

import "fmt"

func logging(applicationName string) {
	fmt.Println("Application: ", applicationName, " has completed")
}

func executeTask(value int) {
	defer logging("Cool APP") // by using defer, the function still executed even the executeTask func was error
	result := value / 2
	fmt.Println("result ", result)
}

func main() {
	executeTask(0)
}

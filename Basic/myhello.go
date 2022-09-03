package main

import "fmt"

func main() {
	fmt.Println("Hi Everyone")

	foo()

	fmt.Println("this is loop or iterative")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("this is foo")
}

func bar() {
	fmt.Println("Exit programm")
}

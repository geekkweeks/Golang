package main

import (
	"fmt"
)

func main() {
	score := "A"

	switch score {
	case "A":
		fmt.Println("Awesomme")
	case "B":
		fmt.Println("Good")
	case "C":
		fmt.Println("Not Bad")
	default:
		fmt.Println("Dont Give Up")
	}

	// Short express
	switch userLength := len("AAZXDF"); userLength < 5 {
	case true:
		fmt.Println("ok")
	case false:
		fmt.Println("hmm")
	}

	// Switch without condition
	age := 30
	switch {
	case age > 21:
		fmt.Println("Old")
	case age < 21:
		fmt.Println("Young")
	}

}

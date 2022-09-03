package main

import "fmt"

func main() {
	var score = 24

	if score <= 50 {
		fmt.Println("Bad")
	} else if score > 50 && score <= 75 {
		fmt.Println("not Bad")
	} else {
		fmt.Println("Perfect")
	}

	// short statement (define variable within expression function)
	if totalChars := len("Alfan Zahriyono"); totalChars < 8 {
		fmt.Println("Too short")
	} else {
		fmt.Println("Too long")
	}

}

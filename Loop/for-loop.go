package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Hello - ", counter)
		counter++
	}

	// For with short statement
	for x := 1; x <= 5; x++ {
		fmt.Println("Short :", x)
	}

	slice := []string{"Ander", "Sule", "Azis"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// For range
	// _ to define that the variable is not needed
	for _, value := range slice {
		fmt.Println(value)
	}

	clubDelegate := make(map[string]string)
	clubDelegate["england"] = "arsenal"
	clubDelegate["spain"] = "barcelona"
	clubDelegate["indonesia"] = "Rans FC"

	for key, value := range clubDelegate {
		fmt.Println("Country: ", key, " Club: ", value)
	}
}

package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Alfan",
		"address": "JL Bambu",
		"phone":   "0292022",
		"hobby":   "Playing game",
	}

	// length of map data
	fmt.Println(len(person))

	//adding new data
	person["status"] = "single"

	fmt.Println(person)

	//get value by key
	fmt.Println(person["name"])

	//delete key in map
	delete(person, "name")

	fmt.Println(person)
}

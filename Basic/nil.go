/*
	- Nil can be applied in several data type:
		1) interface
		2) function
		3) map
		4) slice
		5) Pointer
		6) Channel
*/

package main

import "fmt"

func CountryMap(country string) map[string]string {
	if country == "" {
		return nil
	} else {
		return map[string]string{
			"country": country,
		}
	}
}

func main() {
	var country map[string]string = CountryMap("")

	if country == nil {
		fmt.Println("Empty result")
	} else {
		fmt.Println(country)
	}
}

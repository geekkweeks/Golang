package main

import "fmt"

type isArsenalFans func(string) bool
type clubStadium func(string) string

func enterToArsenalStadium(club string, isArsenalFans isArsenalFans) {
	if isArsenalFans(club) {
		fmt.Println("you are allowed to enter emirate stadium")
	} else {
		fmt.Println(club + " fans are not allowed to enter emirate stadium")
	}
}

func fansFilter(club string) bool {
	if club == "arsenal" {
		return true
	}
	return false
}

func showStadium(name string, clubStadium clubStadium) {
	fmt.Println(clubStadium(name))
}

func main() {
	enterToArsenalStadium("arsenal", fansFilter)
	enterToArsenalStadium("MU", fansFilter)
	enterToArsenalStadium("Liverpool", fansFilter)

	// other anonymous function implementation
	stadium := func(name string) string {
		return "the stadium is " + name
	}

	showStadium("Emirates", stadium)
	showStadium("Senayan", stadium)

	// define function as parameter
	showStadium("Wembley", func(name string) string {
		return "now your stadium is " + name
	})
}

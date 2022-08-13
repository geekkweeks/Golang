package main

import "fmt"

type MoreFilter func(string, string, string) string

func sentenceWithFilter(name string, filter func(string) string) {
	nameWithFilter := filter(name)
	fmt.Println("Hai ", nameWithFilter)
}

func sentenceWithAdvanceFilter(fullName string, address string, phoneNumber string, advFilter MoreFilter) {
	advanceFilter := advFilter(fullName, address, phoneNumber)
	fmt.Println("Hai ", advanceFilter)
}

func advanceFilter(fullName string, address string, phoneNumber string) string {
	fullWorld := ""
	if fullName == "Alfan" {
		fullWorld += "Hei " + fullName
	}
	if address == "Kecapi" {
		fullWorld += "Alamat " + address
	}
	if phoneNumber == "09839" {
		fullWorld += "tlp " + phoneNumber
	}
	return fullWorld
}

func wordFilter(name string) string {
	if name == "Dog" {
		return "..."
	}
	return name
}

func main() {
	sentenceWithFilter("alf", wordFilter)
	sentenceWithFilter("Dog", wordFilter)
	sentenceWithAdvanceFilter("Alfan", "Kecapi", "09839", advanceFilter)
	sentenceWithAdvanceFilter("Alfan", "Kedondong", "09839", advanceFilter)
}

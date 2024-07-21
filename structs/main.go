package main

import "fmt"

type contactInfo struct {
	email       string
	phoneNumber int
}

type person struct {
	firstName  string
	lastName   string
	isVerified bool
	contact    contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func main() {
	var alex person
	alex.firstName = "alex"
	alex.lastName = "Anderson"
	alex.contact = contactInfo{
		email:       "alex@gmail.com",
		phoneNumber: 9089786756,
	}
	alex.print()
}

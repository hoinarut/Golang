package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func main() {
	claudia := person{
		firstName: "Claudia",
		lastName:  "Dabija",
		contactInfo: contactInfo{
			email:   "claudia_ics@yahoo.com",
			zipCode: 7800,
		},
	}

	claudia.print()
	claudiaPointer := &claudia

	claudiaPointer.updateName("Claudia_ics")
	claudia.print()
}

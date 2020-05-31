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
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func main() {

	ambarish := person{
		firstName: "Ambarish",
		lastName:  "Pande",
		contactInfo: contactInfo{
			email:   "ambarish@gmail.com",
			zipCode: 444005,
		},
	}
	ambarish.updateName("Ambya")
	ambarish.print()

}

package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	rahim := person {
		firstName: "Rahim",
		lastName: "Shah",
		contactInfo: contactInfo{
			email: "rahim@email.com",
			zipCode: 95588,
		},
	}
	rahim.updateName("Rahi")
	rahim.print()
	
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}


func (p person) print() {
	fmt.Printf("%+v", p)
}
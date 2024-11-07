package main

import "fmt"

type suit []string

type person struct {
	firstname string
	lastname  string
}

func main() {
	suits := suit{"ACE", "SPADE", "CLUB", "DIAMOND"}
	suits.changeSuit()
	fmt.Printf("suits: %v\n", suits)

	shubhu := person{firstname: "Shubham", lastname: "Pandey"}

	shubhu.changeFirstname()
	fmt.Println(shubhu)
}

// Receiving as reference
func (s *suit) changeSuit() {
	(*s)[0] = "NA"
}

// Receiving as value
func (p person) changeFirstname() {
	p.firstname = "Vikas"
}

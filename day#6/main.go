package main

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	var pandeyJi person
	pandeyContact := contactInfo{email: "shubham.pandey.workconnect@gmail.com", zip: 110059}

	pandeyJi.firstName = "Shubham"
	pandeyJi.lastName = "Pandey"
	pandeyJi.contact = pandeyContact

	pandeyJi.print()
}

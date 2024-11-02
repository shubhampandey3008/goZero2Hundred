package main

import "fmt"

func main() {
	card := returnCard()
	fmt.Println(card)
}

func returnCard() string {
	return "Six of Diamonds"
}

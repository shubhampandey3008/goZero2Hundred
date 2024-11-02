package main

import "fmt"

func main() {

	cards := getCards()
	cards = append(cards, "Three of Clubs")

	for i, card := range getCards() {
		fmt.Println(i, card)
	}
}

func getCards() []string {
	cards := []string{"Ace of Spade", "Five of Hearts", "Seven of Spade"}
	return cards
}

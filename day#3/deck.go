package main

import "fmt"

type deck []string

func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	suits := []string{"Hearts", "Clubs", "Diamonds", "Spades"}
	values := []string{"Ace", "One", "Two", "Three"}

	for _, suit := range suits {

		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, no int) (deck, deck) {
	return d[:no], d[no:]
}

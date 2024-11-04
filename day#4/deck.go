package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) writeToFile() error {
	data := d.toString()
	return os.WriteFile("deckOfCards.txt", []byte(data), 0666)
}

func newDeckFromFile(filename string) deck {
	sob, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	strDeck := string(sob)

	return strings.Split(strDeck, ",")

}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for index := range d {
		newIndex := r.Intn(len(d) - 1)
		d[index], d[newIndex] = d[newIndex], d[index]
	}

}

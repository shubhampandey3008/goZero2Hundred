package main

func main() {
	cards := newDeck()

	cards.writeToFile("testFile")
	// cards := newDeckFromFile("deckOfCards.txt")

	// cards.print()
	// cards.shuffle()
	cards.print()
}

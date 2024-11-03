package main

func main() {
	cards := newDeck()

	cardsPlaced, cardsRemaining := deal(cards, 2)

	cardsPlaced.print()
	cardsRemaining.print()

	cards.print()
}

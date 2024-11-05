package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	cards := newDeck()

	if len(cards) != 16 {
		t.Errorf("Got %v cards instead of 16", len(cards))
	}

	if cards[0] != "Ace of Hearts" {
		t.Errorf("Wrong first card")
	}
}

func TestWriteToFileAndNewDeckFromFile(t *testing.T) {

	os.Remove("_deckFile")
	cards := newDeck()
	cards.writeToFile("_deckFile")
	loadedDeck := newDeckFromFile("_deckFile")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards , got %v", len(loadedDeck))
	}

	os.Remove("_deckFile")
}

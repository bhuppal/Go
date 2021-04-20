package main

import (
	"os"
	"testing")


func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Aces of Spades" {
		t.Errorf("Expected first card on deck should be Aces of Spades but got %v", d[0])
	}

	if d[len(d) -1 ] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d) -1 ])	
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting.txt")

	deck := newDeck()
	deck.saveToFile("_deckTesting.txt")

	loadedDeck := newDeckFromFile("_deckTesting.txt")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_deckTesting.txt")
}

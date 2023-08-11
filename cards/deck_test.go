package main

import (
	"os"
	"testing"
)

// The 't' is our test handler, if anything goes wrong with our tests, we'll tell to this value
// func TestXxx(*testing.T)
func TestNewDeck(t *testing.T) {
	d := newDeck()

	// Code to make sure that a deck is created with x number of cards
	// We have 4 suits and 4 values, we expect our deck to have 16 cards inside
	if len(d) != 16 {
		// Remember to use the interpolation '%v'
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	// Code to make sure that the first card is an Ace of Spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	// Code to make sure that the last card is a Four of Clubs
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

// Testing saveToDeck and newDeckFromFile
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// Delete any files in current working directory with the name "_decktesting"
	// func Remove(name string) error - Remove removes the named file or (empty) directory. If there is an error, it will be of type *PathError.
	os.Remove("_decktesting")

	// Create a new deck;
	deck := newDeck()

	// Save to file "_decktesting"
	deck.saveToFile("_decktesting")

	// Load from file;
	loadedDeck := newDeckFromFile("_decktesting")

	// Assert deck length
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in the deck, but got %v", len(loadedDeck))
	}

	// Delete any files in current working directory with the name "_decktesting"
	os.Remove("_decktesting")
}

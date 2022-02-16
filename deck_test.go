package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Error("Expected deck length of 52, but got ", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Error("Expected aces of spades as first card but got ", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Error("Expected king of clubs as last card but got ", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	ld := newDeckFromFile("_decktesting")

	if len(ld) != 52 {
		t.Error("Expected 52 card deck, got ", len(ld))
	}
}

package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Error("Expected 16, got", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Error("Expected Ace of Spades, got", d[0])
	}
}

func TestSaveToFileAndLoadFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := loadFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Error("Expected 16, got", len(loadedDeck))
	}

	os.Remove("_decktesting")
}

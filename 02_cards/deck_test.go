package main

import (
	"os"
	"testing"
)

// newDeck
func TestNewDeck(t *testing.T) {
	// length
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck lenght of 16, but got %v", len(d))
	}

	// first item
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected deck to start at Ace of Spades, but got %v", d[0])
	}

	// last item
	// course did d[len(d)-1] but I prefer to minimize logic in tests
	if d[15] != "Four of Hearts" {
		t.Errorf("Expected deck to start at Ace of Spades, but got %v", d[0])
	}
}

func TestSaveAndLoadDeck(t *testing.T) {
	testFileName := "_decktesting"
	os.Remove(testFileName)
	d := newDeck()
	d.save(testFileName)
	loadedDeck := loadDeck(testFileName)
	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck lenght of 16, but got %v", len(d))
	}
	os.Remove(testFileName)

}

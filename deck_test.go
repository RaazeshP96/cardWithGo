package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 12 {
		t.Errorf("Expected deck length of 12, but got %v", len(d))

	}
	if d[len(d)-1] != "Cubeofthree" {
		t.Errorf("Expected  last card be cube of three, but got %v", d[len(d)-1])
	}
	if d[0] != "Sphareoface" {
		t.Errorf("Expected  first card be sphare of ace, but got %v", d[0])
	}
}
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 12 {
		t.Errorf("Expected deck length of 12, but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}

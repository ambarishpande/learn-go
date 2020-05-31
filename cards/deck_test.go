package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expectedLength := 52
	if len(d) != expectedLength {
		t.Errorf("Expected length is %v but got %v", expectedLength, len(d))
	}

	expectedFirstCard := "Ace of Hearts"
	if d[0] != expectedFirstCard {
		t.Errorf("Expected first card as %v but got %v", expectedFirstCard, d[0])
	}

	expectedLastCard := "King of Cloves"
	if d[len(d)-1] != expectedLastCard {
		t.Errorf("Expected last card as %v but got %v", expectedLastCard, d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"

	os.Remove(filename)
	d := newDeck()
	d.saveToFile(filename)
	loadedDeck, _ := newDeckFromFile(filename)
	expectedLength := 52
	if len(loadedDeck) != expectedLength {
		t.Errorf("Expected %v cards in deck but found %v", expectedLength, len(loadedDeck))
	}
	os.Remove(filename)

}

package main

import (
	"os"
	"testing"
)

func TestNewDec(t *testing.T) {

	d := newDeck()
	deckLength := len(d)
	expectedFirstElement := "Ace of Spades"
	expectedLastElement := "Four of Cloves"

	if deckLength != 16 {
		t.Errorf("Expected  deck length of 16, but got %v", deckLength)
	}

	if d[0] != expectedFirstElement {
		t.Errorf("Expected first element must be %s", expectedFirstElement)
	}

	actualLastElement := d[deckLength-1]
	if actualLastElement != expectedLastElement {
		t.Errorf("Expected last element is not %v. Actual: %v", expectedLastElement, actualLastElement)
	}

}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	testFile := "_decktesting"

	os.Remove(testFile)

	d := newDeck()
	d.saveToFile(testFile)

	loadedDeck := newDeckFromFile(testFile)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove(testFile)
}

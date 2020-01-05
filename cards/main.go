package main

import "fmt"

func main() {
	// cards := newDeck()
	// cards.saveToFile("hand_cards.csv")
	newCards := newDeckFromFile("hand_cards.csv")
	newCards.print()
	fmt.Println("----------")
	shuffledCards := newCards.shuffle()
	shuffledCards.print()
}

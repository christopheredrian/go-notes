package main

import "fmt"

// Create a new type 'deck': string[]
type deck []string

// New Deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Cloves"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}

// Print function
func (d deck) print() { // d - go idiom to call the receiver as one / two letters
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Deal function - get n cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Add Card

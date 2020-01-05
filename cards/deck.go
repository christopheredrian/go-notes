package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

// To string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Write to File
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666) // 0666 - anyone can read / write the file
}

// New deck from file
func newDeckFromFile(fileName string) deck {
	byteSlices, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(byteSlices), ",")
	return deck(s)
}

// Shuffle
func (d deck) shuffle() deck {

	rand.Seed(time.Now().UnixNano()) // by default go uses a static seed
	maxNum := len(d) - 1

	for i, _ := range d {
		newPosition := rand.Intn(maxNum)
		d[i], d[newPosition] = d[newPosition], d[i]
	}

	return d
}

// Add Card

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// It store the deck of cards as slice of strings
type deck []string

// Returns a new deck of cards
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "One", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Prints the deck of cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Deals the given number of cards from the given deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Encodes the given deck of cards into a string
func (d deck) encode() string {
	encoding := strings.Join(d, ",")
	return encoding
}

// Save the given deck of cards into given file
func (d deck) saveToFile(fileName string) {
	ioutil.WriteFile(fileName, []byte(d.encode()), 0666)
}

package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

// Shuffles the specified deck of cards
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d))
		d[i], d[newPosition] = d[newPosition], d[i]
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

// Decode the encoded deck of cards
func decode(encoding string) deck {
	decoding := strings.Split(encoding, ",")
	return deck(decoding)
}

// Save the given deck of cards into given file
func (d deck) saveToFile(fileName string) {
	ioutil.WriteFile(fileName, []byte(d.encode()), 0666)
}

// Loads the deck of cards from the given file
func loadFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return decode(string(bs))
}

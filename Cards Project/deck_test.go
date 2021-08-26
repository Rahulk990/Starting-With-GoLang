package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()

	if len(cards) != 20 {
		t.Errorf("Expected 20 cards, got %v", len(cards))
	}

	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, got %v", cards[0])
	}

	if cards[len(cards)-1] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs, got %v", cards[len(cards)-1])
	}
}

func TestSavetoFileAndLoadFromFile(t *testing.T) {
	fileName := "_decktesting"
	os.Remove(fileName)

	cards := newDeck()
	cards.saveToFile(fileName)

	loadedCards := loadFromFile(fileName)
	if len(loadedCards) != 200 {
		t.Errorf("Expected 20 loadedCards, got %v", len(loadedCards))
	}

	os.Remove(fileName)
}

package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T){

	d := newDeck()

	if len(d) != 16{
		t.Errorf("Expected 16 but got %v", len(d))
	}

	if d[0] != "Ace of Diamonds"{
		t.Errorf("Expected Ace of Diamonds but %v", d[0])
	}

	if d[len(d)-1] != "Four of Hearts"{
		t.Errorf("Expected Four of Hearts but %v", d[0])
	}

}

func TestSaveDeckAndNewDeckFromFile(t *testing.T){
	fileName := "_decktesting"

	os.Remove(fileName)
	deck := newDeck()
	deck.saveToFile(fileName)

	deck = newDeckFromFile(fileName)

	if len(deck) != 16gi{
		t.Error("Invalid file")
	}

	os.Remove(fileName)

}
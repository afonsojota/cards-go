package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Diamonds", "Spades", "Clubs", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) toString() string{
	result := []string(d)
	return strings.Join(result, ",")
}

func (d deck) saveToFile(filename string) error{
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck{
	file, err := ioutil.ReadFile(filename)
	if(err != nil){
		fmt.Print("Error: " + err.Error())
		os.Exit(1)
	}
	return deck(strings.Split(string(file),","))
}

func (d deck) shuffle(){
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for item := range d{
		newPosition := r.Intn(len(d) -1)

		d[item], d[newPosition] = d[newPosition], d[item]
	}
}
package main

func main() {
	//cards := newDeck()

	deckFromFile := newDeckFromFile("my_cards")

	deckFromFile.shuffle()
	deckFromFile.print()
}

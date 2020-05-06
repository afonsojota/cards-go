package main

func main() {

	deckFromFile := newDeckFromFile("my_cards")

	deckFromFile.shuffle()
	deckFromFile.print()
	
}

package main


// finished 30. Shuffling a Deck

func main() {
	// cards := newDeck()
	// cards.saveToFile("my_cards.txt")
	cards := newDeckFromFile("my_cards.txt")
	cards.printSliceWithComma()
	cards.shuffle()
	cards.printSliceWithComma()
}
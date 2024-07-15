package main

import "fmt"

// finished 40.

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards.txt")
	cards = newDeckFromFile("my_cards.txt")
	cards.printSliceWithComma()
	cards.shuffle()
	cards.printSliceWithComma()

	hand1, hand2 := deal(cards, 5)

	fmt.Println(hand1, hand2)
}
package main

import "fmt"

// finished 23. Multiple Return Values

func main() {
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)
	hand.printSliceWithComma()
	fmt.Println("-----------------------------------------------------------------")
	remainingCards.printSliceWithComma()
}
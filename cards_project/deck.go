package main

import (
	"encoding/json"
	"fmt"
)

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (this deck) printDeck() {
	for i, card := range this {
		fmt.Println(i, card)
	}
}

func (d deck) printSliceWithComma() {
	f, _ := json.Marshal(d)
	fmt.Println(string(f))
}
package main

import "fmt"

//! 'deck' type of strings array
type deck []string

//! create a new deck
func newDeck() deck {
	cards := deck{}

	cardsSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardsValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardsSuits {
		for _, value := range cardsValues {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards
}

//! print all the cards in the deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

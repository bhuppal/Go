package  main

import "fmt"

//Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := deck{"Spades", "Hearts", "Diamonds"}
	cardValues := deck{"Aces", "Two", "Three", "Four"}
	for _, suits := range cardSuits {
		for _, values := range cardValues {
			cards = append(cards, values + " of " + suits)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range  d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}
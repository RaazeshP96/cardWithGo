package main

import "fmt"

// create a new type of "deck"
// which is a slice of string

type deck []string

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Sphare", "Diamonds", "Hearts", "Cube"}
	cardValues := []string{"ace", "two", "three"}
	for i, cardSuit := range cardSuits {
		for j, cardValue := range cardValues {
			cards = append(cards, cardSuit+"of"+cardValue)
		}
	}
}

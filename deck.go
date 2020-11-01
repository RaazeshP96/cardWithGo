package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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
	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardSuit+"of"+cardValue)
		}
	}
	return cards
}
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

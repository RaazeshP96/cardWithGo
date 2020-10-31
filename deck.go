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

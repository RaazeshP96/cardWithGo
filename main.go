package main

import (
	"fmt"
)

func main() {
	// var card string = "this is card"
	card := newCard()
	fmt.Println(card)
}
func newCard() string {
	return "New card"
}

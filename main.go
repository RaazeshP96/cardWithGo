package main

import (
	"fmt"
)

func main() {
	// var card string = "this is card"
	cards := []string{"Five of diamond", newCard()}
	cards = append(cards, "Nine of Sphare")
	for index, card := range cards {
		fmt.Println(index, card)
	}
}
func newCard() string {
	return "Ace of spread"
}

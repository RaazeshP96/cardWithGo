package main

import (
	"fmt"
)

func main() {
	// var card string = "this is card"
	cards := []string{"Five of diamond", newCard()}
	fmt.Println(cards)
}
func newCard() string {
	return "Ace of spread"
}

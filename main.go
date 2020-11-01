package main

import "fmt"

func main() {
	cards := newDeck()
	// hand, reamainCards := deal(cards, 5)
	// hand.print()
	// reamainCards.print()
	fmt.Println(cards.toString())
}

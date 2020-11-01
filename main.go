package main

func main() {
	cards := newDeck()
	// hand, reamainCards := deal(cards, 5)
	// hand.print()
	// reamainCards.print()
	cards.saveToFile("Mycard")
}

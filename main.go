package main

func main() {
	// var card string = "this is card"
	cards := deck{"Five of diamond", newCard()}
	cards = append(cards, "Nine of Sphare")
	cards.print()
}
func newCard() string {
	return "Ace of spread"
}

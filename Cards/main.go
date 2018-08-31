package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = cards.addCard("Six of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}

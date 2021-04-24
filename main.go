package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "six of spades")

	cards.print()
}

func newCard() string {
	return "Five of diamonds"
}

package main

func main() {
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	//-->converting to byte
	// greeting := "Hi there"
	// fmt.Print([]byte(greeting))

	//-->to save to hard disk
	// cards := newDeck()
	// cards.saveToFile("my_cards")

	//-->open from harddisk
	//cards := newDeckFromFile("my_cards")
	//cards := newDeckFromFile("my")----->to get error
	//cards.print()

	//shuffle-->
	cards := newDeck()
	cards.shuffle()
	cards.print()

}

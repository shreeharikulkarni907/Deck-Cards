package main

func main() {
	// cards := newDickFromfile("my_cards")

	// cards.print()

	cards := newDeck()

	cards.shuffle()

	cards.print()

}

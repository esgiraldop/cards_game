package main

func main() {
	cards := newDeckFromFile("hand.txt")
	cards.print()
	cards.shuffle()
	cards.print()
}

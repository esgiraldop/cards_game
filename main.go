package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()
	handToFile := hand.toString()
	fmt.Println("handToFile: ", handToFile)
}

package main

import "fmt"

func main() {
	cards := newDeckFromFile("hand.txt")
	cards.print()
	cards.shuffle()
	fmt.Println("\n")
	cards.print()
}

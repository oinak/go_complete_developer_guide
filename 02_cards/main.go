package main

import (
	"fmt"
	"log"
)

/*
  Variabe declaration structure:
  explicit:
 	var <name> <type> = <value>
  implicit (type inferred)
	<name> := <vaue>
	':' ony first time (initialization)
 - bool
 - string
 - int
 - float64
*/
func main() {
	// deck initialization
	cards := newDeck()

	// deal
	hand, remainingCards := cards.deal(5)
	// hand, remainingCards := deal(cards, 5)

	fmt.Println("Hand:")
	hand.print()
	fmt.Println("Rest:")
	remainingCards.print()

	// save to file
	err := remainingCards.save("remain.deck")
	if err != nil {
		log.Fatal(err)
	}

	// read from file
	savedCards := loadDeck("remain.deck")
	fmt.Println("From file:")
	savedCards.print()

	// shuffle
	savedCards.shuffle()
	fmt.Println("Shuffled:")
	savedCards.print()
}

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Deck is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	suites := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	values := []string{"Ace", "Two", "Three", "Four"}
	for _, suite := range suites {
		for _, value := range values {
			cards = append(cards, value+" of "+suite)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// func deal(d deck, count int) (deck, deck) {
func (d deck) deal(count int) (deck, deck) {
	return d[:count], d[count:]
}

func (d deck) toString() string {
	// we could do
	// s := []strings(d)
	// to use proper types but it's not needed
	return strings.Join(d, ",")
}

func (d deck) save(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0644)
}

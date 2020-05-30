package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Deck is a slice of strings
type deck []string

const separator = ","

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
	return strings.Join(d, separator)
}

func (d deck) save(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0644)
}

func loadDeck(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option 1: log error and retunr newDeck
		// Option 2: log error and entirely quit program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	ss := strings.Split(string(bs), separator)
	cards := deck(ss)
	return cards
}

func (d deck) shuffle() {
	/*
		r := shuffler()
		for i := range d {
			newPosition := r.Intn(len(d) - 1)
			d[i], d[newPosition] = d[newPosition], d[i]
		}
	*/
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
}

// course does not separate this func but I feel like it
// this is needed to avoid having the same pseudo random shuffling every time
func shuffler() *rand.Rand {
	t := time.Now().UnixNano()
	src := rand.NewSource(t)
	r := rand.New(src)
	return r
}

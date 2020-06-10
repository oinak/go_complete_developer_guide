package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	// very custom logic for generating english greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	// very custom logic for generating spanish greeting
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(eb englishBot) string {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) string {
// 	fmt.Println(sb.getGreeting())
// }

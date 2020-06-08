package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	// var colors = map[string]string
	// colors :=make(map[string]string)

	colors["white"] = "#ffffff"

	// colors[1] = 2 // type error on save

	delete(colors, "green")
	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("%v => %v", color, hex)
	}
}

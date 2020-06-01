package main

import "fmt"

// Create a new project folder to house this new project and create a 'main.go' file inside of it.
// Run your code from the terminal by changing into your project directory then running 'go run main.go'.

// In the main.go file, define a function called 'main'.  Remember that the 'main' function will be called automatically.
func main() {
	// In the main function, create a slice of ints from 0 through 10
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// Iterate through the slice with a for loop.  For each element, check to see if the number is even or odd.
	for i := 0; i < 10; i++ {
		// If the value is even, print out "even".  If it is odd, print out "odd"
		if ints[i]%2 == 0 {
			fmt.Printf("%v is even\n", ints[i])
		} else {
			fmt.Printf("%v is odd\n", ints[i])
		}
	}
}

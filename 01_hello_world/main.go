package main // what package does this file belong to?

import "fmt"

func main() {
	greeter()
}

func greeter() {
	fmt.Println("Hi there!")
}

/*
Types of packages:
  - runnable: produces an executable
  - reusable: code as helpers
  package main creates a runable binary
  'main' is a special name

  main packages need a 'main' function

  Packages from the standard library:
  - https://golang.org/pkg/

  File structure:
  1. package declaration
  2. import statements
  3. function declarations
*/

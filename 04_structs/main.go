package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo // the name defaults to the name of the type if omitted
}

func main() {
	// alex := person{"Alex", "Anderson"} // order os definition
	// alex := person{firstName: "Alex", lastName: "Anderson"} // order os definition
	var alex person         // assign type's 'zero-value' per field
	fmt.Println(alex)       // {  }
	fmt.Printf("%+v", alex) // {firstName: lastName:}

	alex.firstName = "Alexander"
	alex.lastName = "Anderson"
	fmt.Printf("%+v", alex) // {firstName:Alexander lastName:Anderson}

	jim := person{
		firstName: "James",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipcode: 12345,
		},
	}
	// fmt.Printf("%+v", jim) // {firstName:Alexander lastName:Anderson}
	// jim.print()

	jim.updateName("Jimmy")
	jim.print()

	jimPointer := &jim // '&' = address of operator
	jimPointer.reallyUpdateName("Max")
	// jim.reallyUpdateName("Foo") works (it's called shortcut)
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p) // {firstName:Alexander lastName:Anderson}
}

// this only updates the internal copy
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p *person) reallyUpdateName(newFirstName string) {
	// '*' = content of pointer operator
	(*p).firstName = newFirstName
}

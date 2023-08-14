package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo or
	contactInfo
}

/*
// First way to declare a struct
func main() {
	alex := person{"Alex", "Anderson"}
	fmt.Println(alex)
}
*/

/*
// Second way to declare a struct
func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
}
*/

/*
// Third way to declare a struct
// A struct is created and the first and last name fields are automatically assigned the value of empty string "", if int or float, assigned to 0 and
// if a Bool, assigned to false, instead of nil, like other languages.

	func main() {
		var alex person
		// Update the properties or the fields on the struct
		alex.firstName = "Alex"
		alex.lastName = "Anderson"

		fmt.Println(alex)
		// Using Printf, "%+v" will print out all the different field names and their values from alex
		fmt.Printf("%+v", alex)
	}
*/
func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		// contact: contactInfo{ or
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	// fmt.Println(jim)

	/*
		jimPointer := &jim -- $ will turn a value into a pointer
		jimPointer.updateName("Jimmy") // jimPointer - type *person, or a pointer to a person
	*/
	jim.updateName("Jimmy") // jim is a type of person

	jim.print()
}

// Go is a pass by value language. So anytime we pass a value to a function, either as a receiver or as an argument, the data is copied in memory.
// This function, by default is always going to be working on a copy of our data structure.
func (pointerToPerson *person) updateName(newFirstName string) { // *person - type *person, or a pointer to a person, turn the pointer into a value.
	(*pointerToPerson).firstName = newFirstName // *pointerToPerson - memory address
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

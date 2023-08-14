package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	// To declare a new value of type struct when the struct has no fields associated with it, just put empty curly braces in front of the struct.
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// Contains identical logic
// func printGreeting(eb spanishBot) {
// 	fmt.Prinln(eb.getGreeting())
// }

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting (Pretending)
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}

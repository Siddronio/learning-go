package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck' which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	// Return everything from the start of the slice to handSize and everything in the deck from the handSize to the very end.
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// This is going to take our slice of string, join it all into one single string separated by commas
	return strings.Join([]string(d), ",")
}

/*
func WriteFile(name string, data []byte, perm FileMode) error
// WriteFile writes data to the named file, creating it if necessary. If the file does not exist, WriteFile creates it with permissions perm
(before umask); otherwise WriteFile truncates it before writing, without changing permissions. Since WriteFile requires multiple
system calls to complete, a failure mid-operation can leave the file in a partially written state.
*/
func (d deck) saveToFile(filename string) error {
	// Save the deck inside to the hard drive
	return os.WriteFile(filename, []byte(d.toString()), 0666)

	/*
		func saveToFile(filename string,d deck) error{
		    cards := strings.Join(d,",")
		    // fmt.Println(, filename)
		    err := os.WriteFile(filename,[]byte(cards),0666)
		    return err

	*/

}

/*
func ReadFile(name string) ([]byte, error)
// ReadFile reads the named file and returns the contents. A successful call returns err == nil, not err == EOF. Because ReadFile reads
the whole file, it does not treat an EOF from Read as an error to be reported.
*/
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error:", err)

		/*
			func Exit(code int)
			// Exit causes the current program to exit with the given status code. Conventionally, code zero indicates success, non-zero an error. The program terminates immediately; deferred functions are not run.
			For portability, the status code should be in the range [0, 125].
		*/
		os.Exit(1)
	}

	/*
		func Split(s, sep string) []string
		// Split slices s into all substrings separated by sep and returns a slice of the substrings between those separators.
		If s does not contain sep and sep is not empty, Split returns a slice of length 1 whose only element is s.
		If sep is empty, Split splits after each UTF-8 sequence. If both s and sep are empty, Split returns an empty slice.
		It is equivalent to SplitN with a count of -1.
		To split around the first instance of a separator, see Cut.
	*/
	// Now we have a []byte and we want a deck. Variable 's' for slice of strings
	s := strings.Split(string(bs), ",")
	// Turn this slice of string into an actual deck
	return deck(s)
}

/*
	Inside package math/rand/
	func Intn(n int) int
	// Intn returns, as an int, a non-negative pseudo-random number in the half-open interval [0,n) from the default Source. It panics if n <=
*/

func (d deck) shuffle() {

	// In package rand/type Rand we have the ability to specify the seed or the source of randomness, so we can call the 'func New'
	// func New(src Source) *Rand
	// - New returns a new Rand that uses random values from src to generate other random values
	// - A Source represents a source of uniformly-distributed pseudo-random int64 values in the range [0, 1<<63). To make a source, we use the 'func NewSource(seed int64) Source'

	// The last thing, pass some value of type64 to this NewSource. Inside the package time, there's 'func Now() Time' that returns the current local time and inside type Time, we have 'func (t Time) UnixNano() int64' that get a value of type 64.
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		// Inside of the type Rand, we have the 'func (r *Rand) Intn(n int) int'
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]

		/*
			// Get the lenght of a slice
			newPosition := rand.Intn(len(d) - 1) // This rand.Intn is a pseudo random generator. Go by default, uses a random number generator that depends upon some seed value, kind like the source of the randomness or our number generator. We're always going to get the exact same sequence becasuse we always start from the same exact seed value

			// Do the logic to kind swap two elements inside our deck
			d[i], d[newPosition] = d[newPosition], d[i] // But they are shuffling in the exact same order (Needs to check - 11 Aug 2023)
		*/
	}

}

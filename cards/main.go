package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

/*
func main() {
		// Before

			cards := newDeckFromFile("my_cards")
			cards.print()

			-------------------------------------------------

			// Create a Deck and save into a file
			cards := newDeck()
			cards.saveToFile("my_cards")

			-------------------------------------------------

			cards := newDeck()
			hand, remainingCards := deal(cards, 5)
			hand.print()
			remainingCards.print()

			-------------------------------------------------

			cards := []string{"Ace of Diamonds", newCard()}
			cards.print()

			for i, card := range cards {
				fmt.Println(i, card)
			}

		// Example of []byte
			greeting := "Hi there!"
			fmt.Println([]byte(greeting))

}
*/

package main

func main() {
	cards := deck{"Five of Diamonds", newCard()} // declare a slice of strings
	cards = append(cards, "Six of Spades")       // append an element to a slice

	cards.print()
}

func newCard() string {
	return "Ace of Spades"
}

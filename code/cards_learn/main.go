package main

import "fmt"

func main() {
	cards := []string{"Five of Diamonds", newCard()} // declare a slice of strings
	cards = append(cards, "Six of Spades")           // append an element to a slice

	for i, card := range cards { // range over a slice NOTE: use :=
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Ace of Spades"
}

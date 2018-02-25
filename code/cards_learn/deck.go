package main

import (
	"fmt"
)

// create a new type of 'deck'
// which is a slice of string

type deck []string

// now you can replace []string with a new deck type

// define a function with a deck as receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

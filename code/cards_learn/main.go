package main

func main() {
	cards := newDeckFromFile("my_cards")
	cards.print()

	newDeckFromFile("not_existed_file")
}

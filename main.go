package main;

// import "fmt";

func main() {
	var cards = newDeck();
	
	cards.print()
}

func newCard() string {
	return "Five of Diamonds";
}
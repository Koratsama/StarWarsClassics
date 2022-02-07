package deck_test

import (
	"fmt"
	"testing"

	"github.com/Koratsama/StarWarsClassics/deck"
)

func TestSabaccDeckInitializationHappyPath(t *testing.T) {
	var deck = deck.InitializeDeck("Sabacc")

	fmt.Printf("\nAll %v cards have been created!\n", len(deck.Cards))

	if len(deck.Cards) != 62 {
		t.Fail()
	}
}

func TestPazaakDeckInitializationHappyPath(t *testing.T) {
	var deck = deck.InitializeDeck("Pazaak")

	//TODO: implement pazaak deck initialization and change test case
	fmt.Printf("\nDeck size of: %v\nNo cards have been created!\n", len(deck.Cards))

	if len(deck.Cards) > 0 {
		t.Fail()
	}
}

func TestSabaccDeckInitializationInvalidOption(t *testing.T) {
	var deck = deck.InitializeDeck("Sabakk")

	fmt.Printf("\nDeck size of: %v\nNo cards have been created!\n", len(deck.Cards))

	if len(deck.Cards) > 0 {
		t.Fail()
	}
}

/*
compares original deck with shuffled deck until a mismatch.
Technically possible for this test case to fail if the shuffle
result happens to match the initial deck order. However, this
is highly unlikely
*/
func TestShuffleDeck(t *testing.T) {
	var initialDeck = deck.InitializeDeck("Sabacc")
	newDeck := deck.Deck{}
	for i := 0; i < len(initialDeck.Cards); i++ {
		newDeck.Cards = append(newDeck.Cards, deck.Card{Stave: initialDeck.Cards[i].Stave, Value: initialDeck.Cards[i].Value})
	}

	shuffledDeck := deck.ShuffleDeck(initialDeck)
	//shuffledDeck := initialDeck
	j := 0
	for i := 0; i < len(shuffledDeck.Cards)-1; i++ {
		fmt.Printf("\n"+shuffledDeck.Cards[i].Stave+" %v", shuffledDeck.Cards[i].Value)
		fmt.Printf("\n"+newDeck.Cards[i].Stave+" %v\n", newDeck.Cards[i].Value)
		if shuffledDeck.Cards[i].Stave == newDeck.Cards[i].Stave && shuffledDeck.Cards[i].Value == newDeck.Cards[i].Value {
			j++
		} else {
			break
		}
	}
	if j == len(shuffledDeck.Cards)-1 {
		t.Fail()
	}
}

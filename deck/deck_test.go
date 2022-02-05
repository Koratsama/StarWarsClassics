package deck_test

import (
	"deck/deck"
	"fmt"
	"testing"
)

func TestSabaccDeckInitialization(t *testing.T) {
	var deck = deck.InitializeDeck("Sabacc")

	fmt.Printf("\n%v\n", len(deck.Cards))

	if len(deck.Cards) != 62 {
		t.Fail()
	}
}

func TestShuffleDeck(t *testing.T) {
	//compares original deck with shuffled deck until a mismatch
	var initialDeck = deck.InitializeDeck("Sabacc")
	newDeck := deck.Deck{}
	for i := 0; i < len(initialDeck.Cards); i++ {
		newDeck.Cards = append(newDeck.Cards, deck.Card{Stave: initialDeck.Cards[i].Stave, Value: initialDeck.Cards[i].Value})
	}

	shuffledDeck := deck.ShuffleDeck(initialDeck)
	j := 0
	for i := 0; i < len(shuffledDeck.Cards)-1; i++ {
		fmt.Printf("\n"+shuffledDeck.Cards[i].Stave+" %v", shuffledDeck.Cards[i].Value)
		fmt.Printf("\n"+newDeck.Cards[i].Stave+" %v", newDeck.Cards[i].Value)
		if shuffledDeck.Cards[i].Stave == newDeck.Cards[i].Stave && shuffledDeck.Cards[i].Value == newDeck.Cards[i].Value {
			j++
			continue
		} else {
			break
		}
	}
	if j == len(shuffledDeck.Cards)-1 {
		t.Fail()
	}
}

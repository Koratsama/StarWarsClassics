package player_test

import (
	"fmt"
	"testing"

	"github.com/Koratsama/StarWarsClassics/deck"
	"github.com/Koratsama/StarWarsClassics/player"
)

func TestDiscardFirstCard(t *testing.T) {
	//initialize a deck of cards.
	var deck = deck.InitializeDeck("Sabacc")
	//create a test player and deal 2 from the deck.
	var testPlayer = player.Player{}
	testPlayer.Hand = deck.Deal(2)

	//identify the card to discard.
	var cardToDiscard = testPlayer.Hand[0]
	fmt.Printf("\nThe first card in the hand we want to discard: %v\n", cardToDiscard)
	//discard the 1st card in the hand of 2.
	var discarded = testPlayer.Discard(1)
	fmt.Printf("\nThe card we actually discarded: %v\n", discarded)

	//if the test players hand is not 1 after discarding, fail the test.
	if len(testPlayer.Hand) != 1 {
		t.Fail()
	}

	//if the discarded card does not match the expected card, fail the test.
	if cardToDiscard != discarded {
		t.Fail()
	}
}

func TestDiscardLastCard(t *testing.T) {
	//initialize a deck of cards.
	var deck = deck.InitializeDeck("Sabacc")
	//create a test player and deal 2 from the deck.
	var testPlayer = player.Player{}
	testPlayer.Hand = deck.Deal(2)

	//identify the card to discard.
	var cardToDiscard = testPlayer.Hand[len(testPlayer.Hand)-1]
	fmt.Printf("\nThe last card in the hand we want to discard: %v\n", cardToDiscard)
	//discard the 1st card in the hand of 2.
	var discarded = testPlayer.Discard(len(testPlayer.Hand))
	fmt.Printf("\nThe card we actually discarded: %v\n", discarded)

	//if the test players hand is not 1 after discarding, fail the test.
	if len(testPlayer.Hand) != 1 {
		t.Fail()
	}

	//if the discarded card does not match the expected card, fail the test.
	if cardToDiscard != discarded {
		t.Fail()
	}
}

func TestFold(t *testing.T) {
	//initialize a deck of cards.
	var deck = deck.InitializeDeck("Sabacc")
	//create a test player and deal 2 from the deck.
	var testPlayer = player.Player{}
	testPlayer.Hand = deck.Deal(2)

	var foldedHand = deck.Cards

	foldedHand = testPlayer.FoldHand()

	fmt.Printf("\nThe cards folded were: %v\n", foldedHand)
	//if the test players hand is not 1 after discarding, fail the test.
	if len(testPlayer.Hand) != 0 {
		fmt.Printf("\ntest player's hand is not empty: %v\n", foldedHand)
		t.Fail()
	}

	//if the discarded card does not match the expected card, fail the test.
	if len(foldedHand) == 0 {
		fmt.Println("\nfolded hand has no cards.")
		t.Fail()
	}
}

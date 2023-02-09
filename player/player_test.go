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

func TestPureSabacc(t *testing.T) {
	var player player.Player
	player.Name = "Test Player"

	addCardToHand(&player, "Sylop", 0)
	addCardToHand(&player, "Sylop", 0)

	player.UpdateHandValue()

	if player.HandCategory != "Sabacc" || player.HandSubCategory != "Pure Sabacc" {
		fmt.Println("\ntest player's hand is not a Pure Sabacc.")
		t.Fail()
	}
}

func TestFullSabacc(t *testing.T) {
	var player player.Player
	player.Name = "Test Player"

	addCardToHand(&player, "Sylop", 0)
	addCardToHand(&player, "triangle", -10)
	addCardToHand(&player, "triangle", 10)
	addCardToHand(&player, "circle", 10)
	addCardToHand(&player, "square", -10)

	player.UpdateHandValue()

	if player.HandCategory != "Sabacc" || player.HandSubCategory != "Full Sabacc" {
		fmt.Println("\ntest player's hand is not a Full Sabacc.")
		t.Fail()
	}
}

func TestFleet(t *testing.T) {
	var player player.Player
	player.Name = "Test Player"

	addCardToHand(&player, "triangle", -3)
	addCardToHand(&player, "triangle", 3)
	addCardToHand(&player, "Sylop", 0)
	addCardToHand(&player, "circle", 3)
	addCardToHand(&player, "square", -3)

	player.UpdateHandValue()

	if player.HandCategory != "Sabacc" || player.HandSubCategory != "Fleet" {
		fmt.Println("\ntest player's hand is not a Fleet.")
		t.Fail()
	}
}

func TestPrimeSabacc(t *testing.T) {
	var player player.Player
	player.Name = "Test Player"

	addCardToHand(&player, "triangle", -10)
	addCardToHand(&player, "triangle", 10)
	addCardToHand(&player, "Sylop", 0)

	player.UpdateHandValue()

	if player.HandCategory != "Sabacc" || player.HandSubCategory != "Prime Sabacc" {
		fmt.Println("\ntest player's hand is not a Prime Sabacc.")
		t.Fail()
	}
}

func TestYeeHaa(t *testing.T) {
	var player player.Player
	player.Name = "Test Player"

	addCardToHand(&player, "triangle", -4)
	addCardToHand(&player, "square", 4)
	addCardToHand(&player, "Sylop", 0)

	player.UpdateHandValue()

	if player.HandCategory != "Sabacc" || player.HandSubCategory != "Yee-Haa" {
		fmt.Println("\ntest player's hand is not a Yee-Haa.")
		t.Fail()
	}
}

func TestRhylet(t *testing.T) {
	var player player.Player
	player.Name = "Test Player"

	addCardToHand(&player, "triangle", -6)
	addCardToHand(&player, "square", -6)
	addCardToHand(&player, "triangle", 4)
	addCardToHand(&player, "square", 4)
	addCardToHand(&player, "circle", 4)

	player.UpdateHandValue()

	if player.HandCategory != "Sabacc" || player.HandSubCategory != "Rhylet" {
		fmt.Println("\ntest player's hand is not a Rhylet.")
		t.Fail()
	}
}

func TestSquadron(t *testing.T) {
	var player player.Player
	player.Name = "Test Player"

	addCardToHand(&player, "circle", -6)
	addCardToHand(&player, "square", -6)
	addCardToHand(&player, "triangle", 6)
	addCardToHand(&player, "square", 6)

	player.UpdateHandValue()

	if player.HandCategory != "Sabacc" || player.HandSubCategory != "Squadron" {
		fmt.Println("\ntest player's hand is not a Squadron.")
		t.Fail()
	}
}

func TestGeeWhiz(t *testing.T) {
	var player player.Player
	player.Name = "Test Player"

	addCardToHand(&player, "circle", 1)
	addCardToHand(&player, "square", 2)
	addCardToHand(&player, "triangle", 3)
	addCardToHand(&player, "square", 4)
	addCardToHand(&player, "square", -10)

	player.UpdateHandValue()

	if player.HandCategory != "Sabacc" || player.HandSubCategory != "Gee Whiz" {
		fmt.Println("\ntest player's hand is not a Gee Whiz.")
		t.Fail()
	}

	player.FoldHand()
	addCardToHand(&player, "circle", -1)
	addCardToHand(&player, "square", -2)
	addCardToHand(&player, "triangle", -3)
	addCardToHand(&player, "square", -4)
	addCardToHand(&player, "square", 10)

	player.UpdateHandValue()

	if player.HandCategory != "Sabacc" || player.HandSubCategory != "Gee Whiz" {
		fmt.Println("\ntest player's hand is not a Gee Whiz.")
		t.Fail()
	}
}

func TestStraightStaves(t *testing.T) {
	var player player.Player
	player.Name = "Test Player"

	addCardToHand(&player, "circle", 1)
	addCardToHand(&player, "square", -2)
	addCardToHand(&player, "triangle", -3)
	addCardToHand(&player, "square", 4)

	player.UpdateHandValue()

	if player.HandCategory != "Sabacc" || player.HandSubCategory != "Straight Staves" {
		fmt.Println("\ntest player's hand is not a Straight Staves.")
		t.Fail()
	}

	player.FoldHand()
	addCardToHand(&player, "circle", 7)
	addCardToHand(&player, "square", -8)
	addCardToHand(&player, "triangle", -9)
	addCardToHand(&player, "square", 10)

	player.UpdateHandValue()

	if player.HandCategory != "Sabacc" || player.HandSubCategory != "Straight Staves" {
		fmt.Println("\ntest player's hand is not a Straight Staves.")
		t.Fail()
	}
}

func TestBanthasWild(t *testing.T) {
	var player player.Player
	player.Name = "Test Player"

	addCardToHand(&player, "circle", -10)
	addCardToHand(&player, "square", -2)
	addCardToHand(&player, "triangle", 4)
	addCardToHand(&player, "square", 4)
	addCardToHand(&player, "circle", 4)

	player.UpdateHandValue()

	if player.HandCategory != "Sabacc" || player.HandSubCategory != "Banthas Wild" {
		fmt.Println("\ntest player's hand is not a Banthas Wild.")
		t.Fail()
	}

	player.FoldHand()
	addCardToHand(&player, "circle", 3)
	addCardToHand(&player, "square", 3)
	addCardToHand(&player, "triangle", -9)
	addCardToHand(&player, "triangle", 3)

	player.UpdateHandValue()

	if player.HandCategory != "Sabacc" || player.HandSubCategory != "Banthas Wild" {
		fmt.Println("\ntest player's hand is not a Banthas Wild.")
		t.Fail()
	}
}

func TestRuleOfTwo(t *testing.T) {
	var player player.Player
	player.Name = "Test Player"

	addCardToHand(&player, "circle", -10)
	addCardToHand(&player, "square", 10)
	addCardToHand(&player, "triangle", 4)
	addCardToHand(&player, "square", -4)
	addCardToHand(&player, "Sylop", 0)

	player.UpdateHandValue()

	if player.HandCategory != "Sabacc" || player.HandSubCategory != "Rule of Two" {
		fmt.Println("\ntest player's hand is not a Rule of Two.")
		t.Fail()
	}

	player.FoldHand()
	addCardToHand(&player, "circle", -3)
	addCardToHand(&player, "square", 3)
	addCardToHand(&player, "triangle", -9)
	addCardToHand(&player, "triangle", 9)

	player.UpdateHandValue()

	if player.HandCategory != "Sabacc" || player.HandSubCategory != "Rule of Two" {
		fmt.Println("\ntest player's hand is not a Rule of Two.")
		t.Fail()
	}

	player.FoldHand()
	addCardToHand(&player, "circle", -6)
	addCardToHand(&player, "square", -2)
	addCardToHand(&player, "triangle", -2)
	addCardToHand(&player, "triangle", 5)
	addCardToHand(&player, "circle", 5)

	player.UpdateHandValue()

	if player.HandCategory != "Sabacc" || player.HandSubCategory != "Rule of Two" {
		fmt.Println("\ntest player's hand is not a Rule of Two.")
		t.Fail()
	}
}

func addCardToHand(player *player.Player, stave string, value int) {
	var card deck.Card
	card.Stave = stave
	card.Value = value
	player.Hand = append(player.Hand, card)
}

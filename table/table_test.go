package table_test

import (
	"fmt"
	"testing"

	"github.com/Koratsama/StarWarsClassics/deck"
	"github.com/Koratsama/StarWarsClassics/table"
)

func TestCheckPlayerStatus(t *testing.T) {
	var testTable = table.Table{}
	testTable.SeatPlayers()
	testTable.SabaccDeck = deck.ShuffleDeck(deck.InitializeDeck("Sabacc"))
	testTable.DealPlayers()

	testTable.Players[0].Credits = 0

	var testGameOver = testTable.CheckPlayerStatus()

	if testGameOver {
		fmt.Println("CheckPlayerStatus func returned incorrectly.")
		t.Fail()
	}

	if len(testTable.Players) != 5 {
		fmt.Println("CheckPlayerStatus func did not remove a player with zero credits.")
		t.Fail()
	}
}

func TestDealPlayers(t *testing.T) {
	//create a test table. Deal 2 cards to each player. The deck should remain with 50 cards.
	var testTable = table.Table{}
	testTable.SeatPlayers()
	testTable.SabaccDeck = deck.ShuffleDeck(deck.InitializeDeck("Sabacc"))
	testTable.DealPlayers()

	if len(testTable.SabaccDeck.Cards) != 50 {
		fmt.Println("The deck does not have the right amount of cards after dealing.")
		t.Fail()
	}

	for _, player := range testTable.Players {
		if len(player.Hand) != 2 {
			fmt.Printf("\n%v does not have hand size of 2.\n", player.Name)
			t.Fail()
		}
	}
}

func TestInitializeDiscardPile(t *testing.T) {
	//create a test table. Initialize the deck with a discard pile
	var testTable = table.Table{}
	testTable.SabaccDeck = deck.ShuffleDeck(deck.InitializeDeck("Sabacc"))
	testTable.InitializeDiscardPile()

	//check the discard pile has one card.
	if len(testTable.DiscardPile) != 1 {
		fmt.Println("table not initialized with a discard pile.")
		t.Fail()
	}

	//check the deck has 61 cards.
	if len(testTable.SabaccDeck.Cards) != 61 {
		fmt.Println("Card not removed from the deck.")
		t.Fail()
	}
}

func TestSeatPlayers(t *testing.T) {
	//create test table and seat 6 players
	var testTable = table.Table{}
	testTable.SeatPlayers()

	if len(testTable.Players) != 6 {
		fmt.Println("table not initialized with 6 players.")
		t.Fail()
	}

	for _, player := range testTable.Players {
		if player.Credits != 300 {
			fmt.Printf("\n%v does not have 300 credits.\n", player.Name)
			t.Fail()
		}
		if len(player.Hand) != 2 {
			fmt.Printf("\n%v does not have hand size of 2.\n", player.Name)
			t.Fail()
		}
	}
}

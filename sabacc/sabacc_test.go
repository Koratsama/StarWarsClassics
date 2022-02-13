package sabacc_test

import (
	"fmt"
	"testing"

	"github.com/Koratsama/StarWarsClassics/sabacc"
	"github.com/Koratsama/StarWarsClassics/table"
)

func TestSabaccStart(t *testing.T) {

	//TODO: actually test user input for the Sabacc games
	fmt.Printf("test Okay!")
}

func TestSetupTable(t *testing.T) {

	table := table.Table{}
	sabacc.SetupTable(&table)

	if len(table.SabaccDeck.Cards) == 0 {
		fmt.Println("Deck not initialized.")
		t.Fail()
	}

	if len(table.Players) != 6 {
		fmt.Println("Not all players seated.")
		t.Fail()
	}

	if len(table.DiscardPile) != 1 {
		fmt.Println("Discard pile not initialized.")
		t.Fail()
	}
}

func TestGain(t *testing.T) {
	table := table.Table{}
	sabacc.SetupTable(&table)

	var testPlayer = table.Players[0]
	sabacc.Gain(&table, &testPlayer)

	if len(table.SabaccDeck.Cards) != 48 {
		fmt.Println("Gain action did not draw a card from the deck.")
		t.Fail()
	}

	if len(testPlayer.Hand) != 3 && len(table.DiscardPile) != 2 {
		fmt.Println("The test player or the discard pile did not gain a card.")
		t.Fail()
	}
}

func TestDiscard(t *testing.T) {
	table := table.Table{}
	sabacc.SetupTable(&table)

	var testPlayer = table.Players[0]
	sabacc.Discard(&table, &testPlayer)

	if len(table.SabaccDeck.Cards) != 48 {
		fmt.Println("Discard action did not draw a card from the deck.")
		t.Fail()
	}

	if len(testPlayer.Hand) != 2 {
		fmt.Println("The test player did not draw or discard.")
		t.Fail()
	}

	if len(table.DiscardPile) != 2 {
		fmt.Println("A card was not added to the discard pile.")
		t.Fail()
	}
}
func TestSwap(t *testing.T) {
	table := table.Table{}
	sabacc.SetupTable(&table)

	var testPlayer = table.Players[0]
	sabacc.Swap(&table, &testPlayer)

	if len(table.SabaccDeck.Cards) != 49 {
		fmt.Println("Discard action did not draw a card from the deck.")
		t.Fail()
	}

	if len(testPlayer.Hand) != 2 {
		fmt.Println("The test player did not properly swap from the discard pile.")
		t.Fail()
	}

	if len(table.DiscardPile) != 1 {
		fmt.Println("A card was not properly swapped into the discard pile.")
		t.Fail()
	}
}

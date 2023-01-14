package sabacc_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/Koratsama/StarWarsClassics/sabacc"
	"github.com/Koratsama/StarWarsClassics/table"
)

func TestSabaccStart(t *testing.T) {

	//TODO: actually test user input for the Sabacc games
	fmt.Println("test Okay!")
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

	//player must keep the card they draw. so there has to be 3 at this point.
	if len(testPlayer.Hand) != 3 {
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

func TestFold(t *testing.T) {
	table := table.Table{}
	sabacc.SetupTable(&table)

	var testPlayer = table.Players[0]
	fmt.Printf("\n%v's hand is: %v\n", testPlayer.Name, testPlayer.Hand)
	sabacc.Fold(&table, &testPlayer)

	if len(testPlayer.Hand) != 0 {
		fmt.Println("The test player did not fold all their cards.")
		fmt.Printf("\n%v's hand after folding is: %v\n", testPlayer.Name, testPlayer.Hand)
		t.Fail()
	}

	if len(table.DiscardPile) != 3 {
		fmt.Println("player did not fold their cards into the discard pile.")
		fmt.Printf("\nDiscard pile: %v\n", table.DiscardPile)
		t.Fail()
	}
}

func TestBetHappyPath(t *testing.T) {
	content := []byte("30")
	tmpfile, err := ioutil.TempFile("", "tempfile")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	os.Stdin = tmpfile

	table := table.Table{}
	sabacc.SetupTable(&table)

	var testPlayer = table.Players[0]
	var playerCredits = testPlayer.Credits
	fmt.Printf("\n%v's hand is: %v\n", testPlayer.Name, testPlayer.Hand)
	var endBet = sabacc.Bet(&table, &testPlayer)

	if testPlayer.Bet == 0 || testPlayer.Credits != playerCredits-30 {
		fmt.Println("The test player did not bet.")
		fmt.Printf("%v's bet is: %v and their total credits are %v.\n", testPlayer.Name, testPlayer.Bet, testPlayer.Credits)
		t.Fail()
	}

	if endBet == false {
		fmt.Println("Error reading player input.")
		t.Fail()
	}

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}

func TestBetInvalidInput(t *testing.T) {
	content := []byte("thirty")
	tmpfile, err := ioutil.TempFile("", "tempfile")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	os.Stdin = tmpfile

	table := table.Table{}
	sabacc.SetupTable(&table)

	var testPlayer = table.Players[0]
	fmt.Printf("\n%v's hand is: %v\n", testPlayer.Name, testPlayer.Hand)
	var endBet = sabacc.Bet(&table, &testPlayer)

	if endBet == false {
		fmt.Println("Error reading player input. --- EXPECTED")
	}

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}

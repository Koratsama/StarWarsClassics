package sabacc

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Koratsama/StarWarsClassics/deck"
	"github.com/Koratsama/StarWarsClassics/player"
	"github.com/Koratsama/StarWarsClassics/table"
)

func Start() {

	var gameOver bool = false
	table := table.Table{}
	SetupTable(&table)

	for !gameOver {

		//TODO: implement game logic
		for _, player := range table.Players {
			fmt.Printf("\nThe discard pile is: %v", table.DiscardPile)
			fmt.Printf("\n%v's hand is: %v", player.Name, player.Hand)
			Action(&table, &player)
			fmt.Printf("\n%v's hand is: %v", player.Name, player.Hand)
		}

		/* BETTING PHASE
		for _, player := range table.Players {
			Bet(&table, &player)
		}*/

		fmt.Printf("\ndiscard pile is: %v", table.DiscardPile)
		fmt.Printf("\nThere are %v cards left in the deck.", len(table.SabaccDeck.Cards))
		gameOver = true
	}
}

func Action(table *table.Table, player *player.Player) {
	var endAction bool = false
	for !endAction {
		var choice string
		endAction = true
		t := time.Now()
		rand.Seed(int64(t.Nanosecond()))
		fmt.Println("\n1. Gain\n2. Discard\n3. Swap\n4. Stand" +
			"\nPlease select an action:")

		fmt.Scanf("%s\n", &choice)

		switch choice {
		case "1", "Gain", "gain":
			Gain(table, player)
		case "2", "Discard", "discard":
			Discard(table, player)
		case "3", "Swap", "swap":
			Swap(table, player)
		case "4", "Stand", "stand":
			//do nothing
		default:
			fmt.Println("Invalid option. Please choose again.")
			endAction = false
		}
	}

}

func Bet(table *table.Table, player *player.Player) {

	var endBet bool = false
	for !endBet {
		var choice string
		endBet = true

		t := time.Now()
		rand.Seed(int64(t.Nanosecond()))
		fmt.Println("\n1. Bet\n2. Check\n3. Fold" +
			"\nPlease select an action:")

		fmt.Scanf("%s\n", &choice)

		switch choice {
		case "1", "Bet", "bet":

			fmt.Println("\nPlease select an amount to bet:")

			fmt.Scanf("%s\n", &choice)
		case "2", "Check", "check":

		case "3", "Fold", "fold":

		default:
			fmt.Println("Invalid option. Please choose again.")
			endBet = false
		}
	}
}

func SetupTable(table *table.Table) {
	fmt.Println("Setting up a table...")
	time.Sleep(1 * time.Second)
	table.SabaccDeck = deck.ShuffleDeck(deck.InitializeDeck("Sabacc"))
	table.SeatPlayers()
	table.DealPlayers()
	table.InitializeDiscardPile()
}

func Gain(table *table.Table, player *player.Player) {
	player.Hand = append(player.Hand, table.SabaccDeck.Deal(1)...)

	//TODO: don't make this a coin flip
	if rand.Intn(2) == 1 {
		table.DiscardPile = append(table.DiscardPile, player.Discard(len(player.Hand)))
	}
}

func Discard(table *table.Table, player *player.Player) {
	table.DiscardPile = append(table.DiscardPile, player.Discard(rand.Intn(len(player.Hand)-1)+1))

	player.Hand = append(player.Hand, table.SabaccDeck.Deal(1)...)
}

func Swap(table *table.Table, player *player.Player) {
	var swappedCard = table.DiscardPile[len(table.DiscardPile)-1]
	table.DiscardPile = table.DiscardPile[:len(table.DiscardPile)-1]
	table.DiscardPile = append(table.DiscardPile, player.Discard(rand.Intn(len(player.Hand)-1)+1))
	player.Hand = append(player.Hand, swappedCard)
}

package sabacc

import (
	"fmt"
	"time"

	"github.com/Koratsama/StarWarsClassics/deck"
	"github.com/Koratsama/StarWarsClassics/table"
)

func Start() {

	var gameOver bool = false
	fmt.Println("Setting up a table...")
	time.Sleep(3 * time.Second)
	table := table.Table{}
	table.SabaccDeck = deck.ShuffleDeck(deck.InitializeDeck("Sabacc"))
	table.SeatPlayers()
	table.DealPlayers()

	for !gameOver {

		//TODO: implement game logic
		//print out the hands and end the game
		for _, player := range table.Players {
			fmt.Printf("\n%v's hand is: "+player.Hand[0].Stave+" %v, "+
				player.Hand[1].Stave+" %v", player.Name, player.Hand[0].Value, player.Hand[1].Value)
		}

		fmt.Printf("\nThere are %v cards left in the deck.", len(table.SabaccDeck.Cards))
		gameOver = true
	}
}

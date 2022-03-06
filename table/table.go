package table

import (
	"fmt"
	"strconv"

	"github.com/Koratsama/StarWarsClassics/deck"
	"github.com/Koratsama/StarWarsClassics/dice"
	"github.com/Koratsama/StarWarsClassics/player"
)

type Table struct {
	SabaccDeck  deck.Deck
	DiscardPile []deck.Card
	Players     []player.Player
	Dice        dice.Dice
	MainPot     int
	SabaccPot   int
	MaxBet      int
}

/*
Name: CheckPlayerStatus
Purpose: The CheckPlayerStatus function checks if there
are enough players for the game to continue. If a player has
no credits left they will be removed from the table.
Parameters: table - reference to the games table
Returns: bool - indicates if the game is over.
*/
func (re *Table) CheckPlayerStatus() bool {
	var updatedPlayers []player.Player
	for _, player := range re.Players {
		if player.Credits > 0 {
			updatedPlayers = append(updatedPlayers, player)
		} else {
			fmt.Printf("%v is out of credits... and left the table.", player.Name)
		}
	}
	re.Players = updatedPlayers
	return len(re.Players) == 1
}

/*
Name: DealPlayers
Purpose: The DealPlayers function deals 2 cards from the top of the
table deck and sets them to each players initial hand for that round.
Parameters: table - reference to the games table.
*/
func (re *Table) DealPlayers() {
	fmt.Println("\nDealing hands...")
	//time.Sleep(1 * time.Second)
	for _, player := range re.Players {
		var hand = re.SabaccDeck.Deal(2)
		player.Hand[0] = hand[0]
		player.Hand[1] = hand[1]
	}
}

/*
Name: InitializeDiscardPile
Purpose: The InitializeDiscardPile function starts the rounds discard pile
by taking one card off the deck and putting it into the tables discardPile.
Parameters: table - reference to the games table.
*/
func (re *Table) InitializeDiscardPile() {
	//time.Sleep(1 * time.Second)
	re.DiscardPile = append(re.DiscardPile, re.SabaccDeck.Deal(1)...)
	//fmt.Printf("\nStarting a Disard Pile with: %v", re.DiscardPile)
}

/*
Name: SeatPlayers
Purpose: The SeatPlayers function populates 6 players around the table. It also
instantiates their hands and total credits to 300.
Parameters: table - reference to the games table
*/
func (re *Table) SeatPlayers() {
	for i := 1; i < 7; i++ {
		player := player.Player{}
		player.Name = "Player " + strconv.Itoa(i)
		player.Credits = 300
		player.Hand = make([]deck.Card, 2)
		re.Players = append(re.Players, player)
	}
}

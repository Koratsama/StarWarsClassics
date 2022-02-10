package table

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Koratsama/StarWarsClassics/deck"
	"github.com/Koratsama/StarWarsClassics/player"
)

type Table struct {
	SabaccDeck  deck.Deck
	DiscardPile []deck.Card
	Players     []player.Player
	MainPot     int
	SabaccPot   int
}

func (re *Table) SeatPlayers() {
	for i := 1; i < 7; i++ {
		player := player.Player{}
		player.Name = "Player " + strconv.Itoa(i)
		player.Credits = 300
		player.Hand = make([]deck.Card, 2)
		re.Players = append(re.Players, player)
	}
}

func (re *Table) DealPlayers() {
	fmt.Println("\nDealing hands...")
	time.Sleep(3 * time.Second)
	for _, player := range re.Players {
		var hand = re.SabaccDeck.Deal(2)
		player.Hand[0] = hand[0]
		player.Hand[1] = hand[1]
	}
}

func (re *Table) InitializeDiscardPile() {
	time.Sleep(1 * time.Second)
	re.DiscardPile = append(re.DiscardPile, re.SabaccDeck.Deal(1)...)
	//fmt.Printf("\nStarting a Disard Pile with: %v", re.DiscardPile)
}

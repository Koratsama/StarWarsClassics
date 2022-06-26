package sabacc

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Koratsama/StarWarsClassics/deck"
	"github.com/Koratsama/StarWarsClassics/player"
	"github.com/Koratsama/StarWarsClassics/table"
)

/*
Name: Start
Purpose: Initiates a game of Sabacc. When the game is over the user
should be brought back to the menu selection.
Parameters: None
*/
func Start() {

	var gameOver bool = false
	table := table.Table{}
	SetupTable(&table)

	for !gameOver {

		//implement round 1
		RoundOne(&table)

		fmt.Println("Round 1 complete!")
		//implement round 2

		//implement round 3

		fmt.Printf("\ndiscard pile is: %v", table.DiscardPile)
		fmt.Printf("\nThere are %v cards left in the deck.", len(table.SabaccDeck.Cards))
		gameOver = true
	}
}

/*
Name: SetupTable
Purpose: Initiates a game of Sabacc.
Parameters: None
*/
func SetupTable(table *table.Table) {
	fmt.Println("Setting up a table...")
	//time.Sleep(1 * time.Second)
	table.SabaccDeck = deck.ShuffleDeck(deck.InitializeDeck("Sabacc"))
	table.Dice.CreateSpikeDice()
	table.SeatPlayers()
	table.DealPlayers()
	table.InitializeDiscardPile()
}

/*
Name: RoundOne
Purpose: This function executes all the steps required for the first round of a game
of sabacc. each player will get prompted to take actions and bets. and then proceed
to the next round.
Parameters: table - reference to the current table.
*/
func RoundOne(table *table.Table) {

	var endRound bool = false

	for i := 0; i < len(table.Players); i++ {
		fmt.Printf("\nThe discard pile [%v] is: %v", len(table.DiscardPile), table.DiscardPile)
		fmt.Printf("\n%v's hand is: %v", table.Players[i].Name, table.Players[i].Hand)
		Action(table, &table.Players[i])
		fmt.Printf("\n%v's hand is: %v", table.Players[i].Name, table.Players[i].Hand)
	}

	for !endRound {
		//loop for betting.
		for i := 0; i < len(table.Players); i++ {
			fmt.Printf("\nThe discard pile [%v] is: %v", len(table.DiscardPile), table.DiscardPile)
			fmt.Printf("\n%v's hand is: %v", table.Players[i].Name, table.Players[i].Hand)
			fmt.Printf("\nCurrent bet is: %v", table.MaxBet)
			if len(table.Players[i].Hand) != 0 {
				BetAction(table, &table.Players[i])
			}
		}
		//check that all players have folded/called. else continue betting.
		endRound = endBetting(table)
	}
}

/*
Name: Action
Purpose: This prompts a player to choose an action. In the context of
a Sabacc game, Players have an action round where they will choose to Gain,
Discard, Swap, or Stand.
Parameters: table, player - reference to the current table and player taking action.
*/
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
			Stand(table, player)
		default:
			fmt.Println("Invalid option. Please choose again.")
			endAction = false
		}
	}

}

/*
Name: Gain
Purpose: The Gain function is for the player to draw 1 from the top of
the deck. The user should then be able to decide to keep or immediately discard
the new card into the discard pile.
Parameters: table, player - reference to the current table and player taking action.
*/
func Gain(table *table.Table, player *player.Player) {
	player.Hand = append(player.Hand, table.SabaccDeck.Deal(1)...)

	//TODO: Double check the rule on discarding the draw. I don't know if you're allowed to discard.
	if rand.Intn(2) == 1 {
		table.DiscardPile = append(table.DiscardPile, player.Discard(len(player.Hand)))
	}
}

/*
Name: Discard
Purpose: The Discard function is for the player to discard 1 card from
their hand into the discard pile and draw a new one from the top of the deck.
Parameters: table, player - reference to the current table and player taking action.
*/
func Discard(table *table.Table, player *player.Player) {
	table.DiscardPile = append(table.DiscardPile, player.Discard(rand.Intn(len(player.Hand)-1)+1))
	player.Hand = append(player.Hand, table.SabaccDeck.Deal(1)...)
}

/*
Name: Swap
Purpose: The Swap function is for the player to take the top card
from the discard pile and swap it with an existing card in their hand.
Parameters: table, player - reference to the current table and player taking action.
*/
func Swap(table *table.Table, player *player.Player) {
	var swappedCard = table.DiscardPile[len(table.DiscardPile)-1]
	table.DiscardPile = table.DiscardPile[:len(table.DiscardPile)-1]
	table.DiscardPile = append(table.DiscardPile, player.Discard(rand.Intn(len(player.Hand)-1)+1))
	player.Hand = append(player.Hand, swappedCard)
}

/*
Name: Stand
Purpose: The Stand function is for the player to essentially
take no action and pass the turn to the next player.
Parameters: table, player - reference to the current table and player taking action.
*/
func Stand(table *table.Table, player *player.Player) {
	fmt.Printf("%v stands\n", player.Name)
}

/*
Name: BetAction
Purpose: The BetAction function will prompt the player with choices
for their betting round. This will include actions such as bet, check, or fold.
Parameters: table, player - reference to the current table and player taking action.
*/
func BetAction(table *table.Table, player *player.Player) {

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
			endBet = Bet(table, player)
		case "2", "Check", "check":
			endBet = Check(table, player)
		case "3", "Fold", "fold":
			Fold(table, player)
		default:
			fmt.Println("Invalid option. Please choose again.")
			endBet = false
		}
	}
}

/*
Name: Bet
Purpose: The Bet function allows the player to specify a number of credits to bet.
if the bet is equivalent to the current tables maximum bet, then the player will
have "called". If the players bet is greater than the current tables maximum bet, then
the player has "Raised" and the betting turns should reset after everyone has taken a
betting action.
Parameters: table, player - reference to the current table and player taking action.
Returns: flag to indicate if player betting action is over.
*/
func Bet(table *table.Table, player *player.Player) bool {
	var bet int
	fmt.Println("\nPlease select an amount to bet:")

	_, err := fmt.Scanf("%d", &bet)
	if err != nil {
		//log.Fatalln(err)
		fmt.Printf("Error reading user input... choose again.\n")
		return false
	}

	if bet < table.MaxBet {
		fmt.Printf("%v did not bet the minimum required: %v\n", player.Name, table.MaxBet)
		return false
	}
	if bet == table.MaxBet {
		fmt.Printf("%v called with %v credits\n", player.Name, bet)
	}
	if bet > table.MaxBet {
		fmt.Printf("%v bet %v credits\n", player.Name, bet)
		table.MaxBet = bet
	}
	player.Bet = bet
	return true
}

/*
Name: Check
Purpose: The Check func essentially passes the betting action to the next player.
This can only be done when the players current bet is equal to the tables maximum bet.
Parameters: table, player - reference to the current table and player taking action.
Returns: flag to indicate if player betting action is over.
*/
func Check(table *table.Table, player *player.Player) bool {
	if table.MaxBet == player.Bet {
		fmt.Printf("%v checks\n", player.Name)
		return true
	} else {
		fmt.Printf("%v cannot check\n", player.Name)
		return false
	}
}

/*
Name: Fold
Purpose: The Fold action allows the player to discard their entire hand into the discard
pile. Once this has happened the player should not be included in the current round.
Parameters: table, player - reference to the current table and player taking action.
*/
func Fold(table *table.Table, player *player.Player) {
	//discard all cards in the hand
	fmt.Printf("Player folded: %v\n", player.Hand)
	table.DiscardPile = append(table.DiscardPile, player.FoldHand()...)
}

func endBetting(table *table.Table) bool {
	//check if all players have either folded or bet matches the maximum bet.
	for _, player := range table.Players {
		if !(len(player.Hand) == 0) && player.Bet != table.MaxBet {
			return false
		}
	}
	return true
}

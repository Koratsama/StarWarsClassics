package sabacc

import (
	"fmt"
	"math"
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
		Round(&table)
		//sabacc shift
		SabaccShift(&table)

		fmt.Println("Round 1 complete!")

		//implement round 2
		Round(&table)
		//sabacc shift
		SabaccShift(&table)

		fmt.Println("Round 2 complete!")
		//implement round 3
		Round(&table)
		//sabacc shift
		SabaccShift(&table)

		fmt.Println("Round 3 complete!")
		//calculate winner

		fmt.Printf("\nThere are %v cards left in the deck.", len(table.SabaccDeck.Cards))
		//show hands
		//decide winner
		var winner player.Player = table.Players[0]
		for _, player := range table.Players {
			if len(player.Hand) >= 2 {
				if int(math.Abs(float64(player.HandValue))) == int(math.Abs(float64(winner.HandValue))) {
					//figure out who wins with hand rules
					winner = decideWinner(winner, player)
				} else if int(math.Abs(float64(player.HandValue))) < int(math.Abs(float64(winner.HandValue))) {
					winner = player
				}
			}
			if len(player.Hand) != 0 {
				fmt.Printf("\n%v's hand total is: %v - %v", player.Name, player.HandValue, player.Hand)
			}
			player.FoldHand()
		}
		fmt.Printf("\n%v wins!! hand total is: %v - %v", winner.Name, winner.HandValue, winner.Hand)
		//TODO: award pot to the winner

		gameOver = true
	}
}

/*
Name: SetupTable
Purpose: Initiates a game of Sabacc.
Parameters: table
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
Name: SetupNewDeck
Purpose: keeps the current state of the given table but creates a new deck and
re-deals new hands to each player still in the game.
Parameters: table

func SetupNewDeck(table *table.Table) {
	fmt.Println("Dealing new hands...")
	//time.Sleep(1 * time.Second)
	//use new deck.
	table.SabaccDeck = deck.ShuffleDeck(deck.InitializeDeck("Sabacc"))
	table.DealPlayers()
	table.InitializeDiscardPile()
}*/

/*
Name: Round
Purpose: This function executes all the steps required for the first round of a game
of sabacc. each player will get prompted to take actions and bets. and then proceed
to the next round.
Parameters: table - reference to the current table.
*/
func Round(table *table.Table) {

	var endRound bool = false

	for i := 0; i < len(table.Players); i++ {
		if len(table.Players[i].Hand) != 0 {
			fmt.Printf("\nThe discard pile [%v] is: %v", len(table.DiscardPile), table.DiscardPile[len(table.DiscardPile)-1])
			fmt.Printf("\n%v's hand is: %v", table.Players[i].Name, table.Players[i].Hand)
			Action(table, &table.Players[i])
			fmt.Printf("\n%v's hand is: %v", table.Players[i].Name, table.Players[i].Hand)
		}
	}

	for !endRound {
		//loop for betting.
		for i := 0; i < len(table.Players); i++ {
			if len(table.Players[i].Hand) != 0 {
				if table.Players[i].AllIn {
					fmt.Printf("\n%v is all in already.\n", table.Players[i].Name)
					continue
				}
				fmt.Printf("\nThe discard pile [%v] is: %v", len(table.DiscardPile), table.DiscardPile[len(table.DiscardPile)-1])
				fmt.Printf("\n%v's hand is: %v", table.Players[i].Name, table.Players[i].Hand)
				fmt.Printf("\nCurrent bet is: %v", table.MaxBet)
				BetAction(table, &table.Players[i])
			}
		}
		//check that all players have folded/called. else continue betting.
		endRound = endBetting(table)
	}
}

/*
Name: SabaccShift
Purpose: This function should be executed after each round. the table should have
the spike dice rolled. If there is a sabacc shift,
Parameters: table - reference to the current table.
*/
func SabaccShift(table *table.Table) {

	var shift bool = table.Dice.RollSpikeDice()

	if shift {
		fmt.Println("There's been a shift!")
		for i := 0; i < len(table.Players); i++ {
			var handSize int = len(table.Players[i].Hand)
			if handSize != 0 {
				table.DiscardPile = append(table.DiscardPile, table.Players[i].Discard(handSize))
				//table.Players[i].Discard(handSize)
				table.Players[i].Hand = table.SabaccDeck.Deal(handSize)
				table.Players[i].UpdateHandValue()
			}
		}
	} else {
		fmt.Println("No Sabacc Shift this round.")
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
the deck. The user has to keep this card.
Parameters: table, player - reference to the current table and player taking action.
*/
func Gain(table *table.Table, player *player.Player) {
	player.Hand = append(player.Hand, table.SabaccDeck.Deal(1)...)
	player.UpdateHandValue()
}

/*
Name: Discard
Purpose: The Discard function is for the player to discard 1 card from
their hand into the discard pile and draw a new one from the top of the deck.
Parameters: table, player - reference to the current table and player taking action.
*/
func Discard(table *table.Table, player *player.Player) {
	//TODO: choose which card you want to discard.
	table.DiscardPile = append(table.DiscardPile, player.Discard(rand.Intn(len(player.Hand)-1)+1))
	player.Hand = append(player.Hand, table.SabaccDeck.Deal(1)...)
	player.UpdateHandValue()
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
	player.UpdateHandValue()
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

	_, err := fmt.Scanf("%d\n", &bet)
	if err != nil {
		//log.Fatalln(err)
		fmt.Printf("Error reading user input... choose again.\n")
		return false
	}

	if bet < table.MaxBet {
		if bet > player.Bet+player.Credits {
			fmt.Printf("%v does not have enough credits to bet %v. Total Credits: %v\n", player.Name, bet, player.Bet+player.Credits)
			return false
		} else if bet < player.Bet+player.Credits {
			fmt.Printf("%v did not bet the minimum required: %v\n", player.Name, table.MaxBet)
			return false
		} else if bet == player.Bet+player.Credits {
			fmt.Printf("%v called with %v credits\n", player.Name, bet)
			fmt.Printf("%v is all in!\n", player.Name)
			player.AllIn = true
		}
	}
	if bet == table.MaxBet {
		if bet > player.Bet+player.Credits {
			fmt.Printf("%v does not have enough credits to bet %v. Total Credits: %v\n", player.Name, bet, player.Bet+player.Credits)
			return false
		} else {
			fmt.Printf("%v called with %v credits\n", player.Name, bet)
			if bet == player.Bet+player.Credits {
				fmt.Printf("%v is all in!\n", player.Name)
				player.AllIn = true
			}
		}
		//if players bet is equal to max bet subtract the difference between the
		//called bet and the players current bet
		player.Credits = player.Credits - (bet - player.Bet)
	}
	if bet > table.MaxBet {
		if bet > player.Bet+player.Credits {
			fmt.Printf("%v does not have enough credits to bet %v. Total Credits: %v\n", player.Name, bet, player.Bet+player.Credits)
			return false
		}
		fmt.Printf("%v bet %v credits\n", player.Name, bet)
		if bet == player.Bet+player.Credits {
			fmt.Printf("%v is all in!\n", player.Name)
			player.AllIn = true
		}
		table.MaxBet = bet
		//subtract the difference between the new max bet and the players current bet.
		player.Credits = player.Credits - (bet - player.Bet)
	}
	player.Bet = bet
	table.MainPot += bet
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

	for i := 0; i < len(table.Players); i++ {
		table.Players[i].Bet = 0
	}
	table.MaxBet = 0
	return true
}

func decideWinner(currentWinner player.Player, nextPlayer player.Player) player.Player {

	//TODO: logic to decide real winner if there is a tie
	return currentWinner
}

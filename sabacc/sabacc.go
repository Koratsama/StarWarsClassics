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

		for _, player := range table.Players {
			fmt.Printf("\nThe discard pile is: %v", table.DiscardPile)
			fmt.Printf("\n%v's hand is: %v", player.Name, player.Hand)
			fmt.Printf("\nCurrent bet is: %v", table.MaxBet)
			if len(player.Hand) != 0 {
				BetAction(&table, &player)
			}
		}

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

func SetupTable(table *table.Table) {
	fmt.Println("Setting up a table...")
	//time.Sleep(1 * time.Second)
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

func Check(table *table.Table, player *player.Player) bool {
	if table.MaxBet == player.Bet {
		fmt.Printf("%v checks\n", player.Name)
		return true
	} else {
		fmt.Printf("%v cannot check\n", player.Name)
		return false
	}
}

func Fold(table *table.Table, player *player.Player) {
	//discard all cards in the hand
	fmt.Printf("Player folded: %v\n", player.Hand)
	table.DiscardPile = append(table.DiscardPile, player.FoldHand()...)
}

package player

import "github.com/Koratsama/StarWarsClassics/deck"

type Player struct {
	Name      string
	Hand      []deck.Card
	Credits   int
	Position  int
	Bet       int
	AllIn     bool
	HandValue int
}

/*
Name: Discard
Purpose: This allows the player to choose which card in their hand
they want to discard. In the context of a Sabacc game, this is part of
the Swap action where a player discards a card and then draws a new one.
Parameters: cardNumber int - the card in the hand to discard.
*/
func (re *Player) Discard(cardNumber int) deck.Card {
	var DiscardCard deck.Card = re.Hand[cardNumber-1]

	if cardNumber == len(re.Hand) {
		re.Hand = re.RemoveLastCard()
	} else {
		re.Hand = append(re.Hand[:cardNumber-1], re.Hand[cardNumber:]...)
	}

	return DiscardCard
}

/*
Name: FoldHand
Purpose: The purpose of this function is to fold the hand of the player.
This will result in the entire hand being returned so that it can be placed
on the top of the discard pile.
Parameters: None
*/
func (re *Player) FoldHand() []deck.Card {
	var discardHand []deck.Card = re.Hand
	re.Hand = make([]deck.Card, 0)
	return discardHand
}

/*
Name: RemoveLastCard
Purpose: The purpose of this function is to remove the last card from a player's hand.
In the context of a sabacc game, this is used when a player gains a card but decides to
discard that drawn card immediately.
Parameters: None
*/
func (re *Player) RemoveLastCard() []deck.Card {
	return re.Hand[:len(re.Hand)-1]
}

/*
Name: UpdateHandValue
Purpose: the purpose of the function is to calculate the players hand value whenever
the player takes an action that changes their hand such as Gain, Discard, Swap.
After each of these actions this function should be called to recalculate the
player's hand value.
Paramets: None
*/
func (re *Player) UpdateHandValue() {
	total := 0
	hand := re.Hand
	for i := range hand {
		total += hand[i].Value
	}
	re.HandValue = total
}

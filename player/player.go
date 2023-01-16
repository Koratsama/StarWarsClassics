package player

import (
	"math"

	"github.com/Koratsama/StarWarsClassics/deck"
)

type Player struct {
	Name                string
	Hand                []deck.Card
	Credits             int
	Position            int
	Bet                 int
	AllIn               bool
	HandValue           int
	HandCategory        string
	HandSubCategory     string
	PositiveCards       int
	HighestPositiveCard deck.Card
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
Parameters: None
*/
func (re *Player) UpdateHandValue() {
	total := 0
	re.PositiveCards = 0
	re.HighestPositiveCard = re.Hand[0]

	hand := re.Hand
	for i := range hand {
		if hand[i].Value > re.HighestPositiveCard.Value {
			re.HighestPositiveCard = hand[i]
		}
		if hand[i].Value > 0 {
			re.PositiveCards += 1
		}
		total += hand[i].Value
	}
	re.HandValue = total

	if total == 0 {
		re.HandCategory = "Sabacc"
		//check what kind of Sabacc
		if isPureSabacc(hand) {
			re.HandSubCategory = "Pure Sabacc"
		} else if isFullSabacc(hand) {
			re.HandSubCategory = "Full Sabacc"
		} else if isFleet(hand) {
			re.HandSubCategory = "Fleet"
		} else if isPrimeSabacc(hand) {
			re.HandSubCategory = "Prime Sabacc"
		} else if isYeeHaa(hand) {
			re.HandSubCategory = "Yee-Haa"
		}
	} else {
		re.HandCategory = "Nulrhek"
	}
}

func isPureSabacc(hand []deck.Card) bool {

	if len(hand) == 2 && hand[0].Value == 0 && hand[1].Value == 0 {
		return true
	} else {
		return false
	}
}

func isFullSabacc(hand []deck.Card) bool {

	positiveTens := 0
	negativeTens := 0
	sylop := 0

	if len(hand) != 5 {
		return false
	} else {
		for i := range hand {
			if hand[i].Value == -10 {
				negativeTens++
			} else if hand[i].Value == 10 {
				positiveTens++
			} else if hand[i].Value == 0 {
				sylop++
			} else {
				return false
			}
		}
	}

	if positiveTens != 2 || negativeTens != 2 || sylop != 1 {
		return false
	} else {
		return true
	}
}

func isFleet(hand []deck.Card) bool {
	//four of a kind with a sylop
	fourOfAKindValue := 0
	fourOfAKind := 0
	sylop := 0

	if len(hand) != 5 {
		return false
	} else {
		for i := range hand {
			if hand[i].Value == 0 {
				sylop++
			} else if hand[i].Value != 0 {
				if fourOfAKindValue == 0 {
					fourOfAKindValue = int(math.Abs(float64(hand[i].Value)))
				} else if fourOfAKindValue != 0 && int(math.Abs(float64(hand[i].Value))) != fourOfAKindValue {
					return false
				}
			} else {
				fourOfAKind++
			}
		}
	}

	if sylop != 1 || fourOfAKind != 4 {
		return false
	} else {
		return true
	}
}

func isPrimeSabacc(hand []deck.Card) bool {

	positiveTens := 0
	negativeTens := 0
	sylop := 0

	if len(hand) != 3 {
		return false
	} else {
		for i := range hand {
			if hand[i].Value == -10 {
				negativeTens++
			} else if hand[i].Value == 10 {
				positiveTens++
			} else if hand[i].Value == 0 {
				sylop++
			} else {
				return false
			}
		}
	}

	if positiveTens != 1 || negativeTens != 1 || sylop != 1 {
		return false
	} else {
		return true
	}
}

func isYeeHaa(hand []deck.Card) bool {
	//pair with a sylop
	pairValue := 0
	pair := 0
	sylop := 0

	if len(hand) != 5 {
		return false
	} else {
		for i := range hand {
			if hand[i].Value == 0 {
				sylop++
			} else if hand[i].Value != 0 {
				if pairValue == 0 {
					pairValue = int(math.Abs(float64(hand[i].Value)))
				} else if pairValue != 0 && int(math.Abs(float64(hand[i].Value))) != pairValue {
					return false
				}
			} else {
				pair++
			}
		}
	}

	if sylop != 1 || pair != 2 {
		return false
	} else {
		return true
	}
}

package player

import (
	"math"
	"math/rand"
	"sort"

	"github.com/Koratsama/StarWarsClassics/deck"
)

type Player struct {
	Name                string
	Hand                []deck.Card
	Credits             int
	Position            int
	Bet                 int
	AllIn               bool
	HasSylop            bool
	HandRank            int
	HandValue           int
	HandCategory        string
	HandSubCategory     string
	PositiveCards       int
	PositiveCardTotal   int
	HighestPositiveCard deck.Card
	Skill    			int
    Charisma 			int
    Risk     			int
    Cheat    			int
    Drunk    			int
    Anger    			int
}

func PreMadeCharacters() []Player {
    return []Player{
            {Name: "Han Solo", Credits: 1000, Hand: make([]deck.Card, 2), Skill: 90, Charisma: 80, Risk: 70, Cheat: 60, Drunk: 20, Anger: 30},
            {Name: "Chewbacca", Credits: 1000, Hand: make([]deck.Card, 2), Skill: 70, Charisma: 50, Risk: 60, Cheat: 10, Drunk: 10, Anger: 40},
            {Name: "Boba Fett", Credits: 1000, Hand: make([]deck.Card, 2), Skill: 85, Charisma: 40, Risk: 80, Cheat: 70, Drunk: 10, Anger: 50},
            {Name: "Jabba the Hutt", Credits: 1000, Hand: make([]deck.Card, 2), Skill: 60, Charisma: 30, Risk: 50, Cheat: 80, Drunk: 50, Anger: 60},
            {Name: "Darth Vader", Credits: 1000, Hand: make([]deck.Card, 2), Skill: 95, Charisma: 70, Risk: 90, Cheat: 50, Drunk: 10, Anger: 80},
            {Name: "Obi-wan Kenobi", Credits: 1000, Hand: make([]deck.Card, 2), Skill: 90, Charisma: 80, Risk: 60, Cheat: 20, Drunk: 10, Anger: 20},
            {Name: "C-3PO", Credits: 1000, Hand: make([]deck.Card, 2), Skill: 50, Charisma: 90, Risk: 30, Cheat: 10, Drunk: 0, Anger: 10},
            {Name: "R2D2", Credits: 1000, Hand: make([]deck.Card, 2), Skill: 70, Charisma: 60, Risk: 50, Cheat: 20, Drunk: 0, Anger: 20},
            {Name: "Bossk", Credits: 1000, Hand: make([]deck.Card, 2), Skill: 80, Charisma: 40, Risk: 70, Cheat: 60, Drunk: 20, Anger: 50},
            {Name: "Lando Calrissian", Credits: 1000, Hand: make([]deck.Card, 2), Skill: 85, Charisma: 90, Risk: 75, Cheat: 50, Drunk: 30, Anger: 40},
        }
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
	re.HandCategory = "Folded"
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
Name: DecideAction
Purpose: The purpose of this function is to simulate a player making a decision based on
their attributes and current hand value. This function will return a string that represents the decision that the
player has made. The decision can be one of the following: Gain, Discard, Swap, Stand, or cheat.
Attributes:
- Skill: The player's skill level. Higher skill means better decision making.
- Charisma: The player's charisma level. Higher Charisma means more chatty.
- Risk: The player's risk level. Higher risk means more unpredictable.
- Cheat: The player's cheat level. Higher cheat means more likely to know another player's hand.
- Drunk: The player's drunk level. Hugher drunk means less likely to make good decisions. can change throughout a game.
- Anger: The player's anger level. Higher anger means more likely to make aggressive decisions.
*/
func (p *Player) DecideAction() string {
    
    decision := "Stand"

    // Calculate a decision score based on attributes and hand value
	//TODO: change this later
    score := p.Skill + p.Charisma - p.Drunk + p.Risk - p.Anger

    // Random factor to simulate unpredictability
    randomFactor := rand.Intn(100)

    // Example logic to decide based on score and random factor
    if score+randomFactor > 150 {
        decision = "Swap"
    } else if score+randomFactor > 100 {
        decision = "Gain"
    } else if score+randomFactor > 50 {
		decision = "Discard"
	}

    // TODO: Implement cheating logic
    // if p.Cheat > 50 && rand.Intn(100) < p.Cheat {
    //     decision = "Cheat"
    // }

    return decision
}

/*
Name: DecideBettingAction
Purpose: The purpose of this function is to simulate a player making a bet decision based on
their attributes and current hand value. This function will return a string that represents the decision that the
player has made. The decision can be one of the following: Bet, Check, Fold.
Attributes:
- Skill: The player's skill level. Higher skill means better decision making.
- Charisma: The player's charisma level. Higher Charisma means more chatty.
- Risk: The player's risk level. Higher risk means more unpredictable.
- Cheat: The player's cheat level. Higher cheat means more likely to know another player's hand.
- Drunk: The player's drunk level. Hugher drunk means less likely to make good decisions. can change throughout a game.
- Anger: The player's anger level. Higher anger means more likely to make aggressive decisions.
*/
func (p *Player) DecideBetAction() string {
    
    decision := "Fold"

    // Calculate a decision score based on attributes and hand value
	//TODO: change this later
    score := p.Skill + p.Charisma - p.Drunk + p.Risk - p.Anger

    // Random factor to simulate unpredictability
    randomFactor := rand.Intn(100)

    // Example logic to decide based on score and random factor
    if score+randomFactor > 150 {
        decision = "Bet"
    } else if score+randomFactor > 100 {
        decision = "Check"
    } 

    return decision
}

/*
Name: DecideBetAmount
Purpose: The purpose of this function is to simulate a player deciding the bet amount based on
their attributes and current hand value. This function will return an integer that represents the bet amount.
Attributes:
- Skill: The player's skill level. Higher skill means better decision making.
- Risk: The player's risk level. Higher risk means more unpredictable.
- Drunk: The player's drunk level. Higher drunk means less likely to make good decisions.
- Anger: The player's anger level. Higher anger means more likely to make aggressive decisions.
- HandValue: The value of the player's current hand.
*/
func (p *Player) DecideBetAmount(currentBet int) int {
    // Base bet amount
    baseBet := 10

    // Calculate a bet multiplier based on attributes and hand value
    multiplier := float64(p.Skill)/100 + float64(p.Risk)/100 - float64(p.Drunk)/200 + float64(p.HandValue)/100 - float64(p.Anger)/200

    // Random factor to simulate unpredictability
    randomFactor := rand.Float64()

    // Calculate the final bet amount
    betAmount := int(float64(baseBet) * multiplier * (1 + randomFactor))

    // Ensure the bet amount is at least the base bet
    if betAmount < baseBet {
        betAmount = baseBet
    }

	if (betAmount < currentBet) {
		betAmount = currentBet
	}
    // Ensure the bet amount does not exceed the player's credits
    if betAmount > p.Credits {
        betAmount = p.Credits
    }

    return betAmount
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
	re.HandRank = 0
	re.PositiveCards = 0
	re.PositiveCardTotal = 0
	re.HighestPositiveCard = re.Hand[0]

	hand := re.Hand
	for i := range hand {
		if hand[i].Value > re.HighestPositiveCard.Value {
			re.HighestPositiveCard = hand[i]
		}
		if hand[i].Value > 0 {
			re.PositiveCards += 1
			re.PositiveCardTotal += hand[i].Value
		}
		if hand[i].Value == 0 {
			re.HasSylop = true
		}
		total += hand[i].Value
	}
	re.HandValue = total

	if total == 0 {
		re.HandCategory = "Sabacc"
		//check what kind of Sabacc
		if isPureSabacc(hand) {
			re.HandSubCategory = "Pure Sabacc"
			re.HandRank = 1
		} else if isFullSabacc(hand) {
			re.HandSubCategory = "Full Sabacc"
			re.HandRank = 2
		} else if isFleet(hand) {
			re.HandSubCategory = "Fleet"
			re.HandRank = 3
		} else if isPrimeSabacc(hand) {
			re.HandSubCategory = "Prime Sabacc"
			re.HandRank = 4
		} else if isYeeHaa(hand) {
			re.HandSubCategory = "Yee-Haa"
			re.HandRank = 5
			//at this point the player doesn't have a sylop
		} else if isRhylet(hand) {
			re.HandSubCategory = "Rhylet"
			re.HandRank = 6
		} else if isSquadron(hand) {
			re.HandSubCategory = "Squadron"
			re.HandRank = 7
		} else if isGeeWhiz(hand) {
			re.HandSubCategory = "Gee Whiz"
			re.HandRank = 8
		} else if isStraightStaves(hand) {
			re.HandSubCategory = "Straight Staves"
			re.HandRank = 9
		} else if isBanthasWild(hand) {
			re.HandSubCategory = "Banthas Wild"
			re.HandRank = 10
		} else if isRuleOfTwo(hand) {
			re.HandSubCategory = "Rule of Two"
			re.HandRank = 11
		} else {
			re.HandSubCategory = "none"
			re.HandRank = 12
		}
	} else {
		re.HandCategory = "Nulrhek"
		re.HandSubCategory = "none"
		re.HandRank = 14
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
					fourOfAKind++
				} else if fourOfAKindValue != 0 && int(math.Abs(float64(hand[i].Value))) != fourOfAKindValue {
					return false
				} else {
					fourOfAKind++
				}
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

	if len(hand) != 3 {
		return false
	} else {
		for i := range hand {
			if hand[i].Value == 0 {
				sylop++
			} else if hand[i].Value != 0 {
				if pairValue == 0 {
					pairValue = int(math.Abs(float64(hand[i].Value)))
					pair++
				} else if pairValue != 0 && int(math.Abs(float64(hand[i].Value))) != pairValue {
					return false
				} else {
					pair++
				}
			}
		}
	}

	if sylop != 1 || pair != 2 {
		return false
	} else {
		return true
	}
}

func isRhylet(hand []deck.Card) bool {
	//four of a kind with a sylop
	threeOfAKindValue := 0
	threeOfAKind := 0
	pairValue := 0
	pair := 0

	if len(hand) != 5 {
		return false
	} else {
		for i := range hand {
			if threeOfAKindValue == 0 {
				threeOfAKindValue = hand[i].Value
				threeOfAKind++
			} else if pairValue == 0 && hand[i].Value != threeOfAKindValue {
				pairValue = hand[i].Value
				pair++
			} else if hand[i].Value != threeOfAKindValue && hand[i].Value != pairValue {
				return false
			} else if hand[i].Value == threeOfAKindValue {
				threeOfAKind++
			} else if hand[i].Value == pairValue {
				pair++
			}
		}
	}

	if (pair == 2 && threeOfAKind == 3) || (pair == 3 && threeOfAKind == 2) {
		return true
	} else {
		return false
	}
}

func isSquadron(hand []deck.Card) bool {
	//four of a kind without a sylop
	fourOfAKindValue := 0
	fourOfAKind := 0

	if len(hand) != 4 {
		return false
	} else {
		for i := range hand {
			if hand[i].Value != 0 {
				if fourOfAKindValue == 0 {
					fourOfAKindValue = int(math.Abs(float64(hand[i].Value)))
					fourOfAKind++
				} else if fourOfAKindValue != 0 && int(math.Abs(float64(hand[i].Value))) != fourOfAKindValue {
					return false
				} else {
					fourOfAKind++
				}
			}
		}
	}

	if fourOfAKind == 4 {
		return true
	} else {
		return false
	}
}

func isGeeWhiz(hand []deck.Card) bool {
	// 1,2,3,4 and -10
	positiveOne := false
	positiveTwo := false
	positiveThree := false
	positiveFour := false
	negativeTen := false
	//-1,-2,-3,-4 and 10
	negativeOne := false
	negativeTwo := false
	negativeThree := false
	negativeFour := false
	positiveTen := false

	if len(hand) != 5 {
		return false
	} else {
		for i := range hand {
			switch hand[i].Value {
			case 1:
				positiveOne = true
			case 2:
				positiveTwo = true
			case 3:
				positiveThree = true
			case 4:
				positiveFour = true
			case -10:
				negativeTen = true
			case -1:
				negativeOne = true
			case -2:
				negativeTwo = true
			case -3:
				negativeThree = true
			case -4:
				negativeFour = true
			case 10:
				positiveTen = true
			default:
				return false
			}
		}
	}

	if (positiveOne && positiveTwo && positiveThree && positiveFour && negativeTen) ||
		(negativeOne && negativeTwo && negativeThree && negativeFour && positiveTen) {
		return true
	} else {
		return false
	}
}

func isStraightStaves(hand []deck.Card) bool {
	//fix this, try sorting instead

	sort.Slice(hand, func(i, j int) bool {
		return int(math.Abs(float64(hand[i].Value))) < int(math.Abs(float64(hand[j].Value)))
	})

	if len(hand) != 4 {
		return false
	} else {
		first := int(math.Abs(float64(hand[0].Value)))
		second := int(math.Abs(float64(hand[1].Value)))
		third := int(math.Abs(float64(hand[2].Value)))
		fourth := int(math.Abs(float64(hand[3].Value)))
		if first == second-1 && second == third-1 && third == fourth-1 {
			return true
		} else {
			return false
		}
	}
}

func isBanthasWild(hand []deck.Card) bool {

	valueMap := make(map[int]int)

	if len(hand) < 4 {
		return false
	} else {
		for i := range hand {
			valueMap[int(math.Abs(float64(hand[i].Value)))]++
		}

		for i := range hand {
			if valueMap[int(math.Abs(float64(hand[i].Value)))] == 3 {
				return true
			}
		}
	}
	return false
}

func isRuleOfTwo(hand []deck.Card) bool {

	valueMap := make(map[int]int)
	numOfPairs := 0

	if len(hand) < 4 {
		return false
	} else {
		for i := range hand {
			valueMap[int(math.Abs(float64(hand[i].Value)))]++
		}

		for i := range valueMap {
			if valueMap[i] == 2 {
				numOfPairs++
			}
		}

		if numOfPairs == 2 {
			return true
		} else {
			return false
		}
	}
}

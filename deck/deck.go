package deck

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
}

type Card struct {
	Stave string
	Value int
}

/**
Name: Deal
Purpose: Returns the specified number of cards from the top of the deck.
In the context of any card game this is useful for dealing out a number
of cards.
Parameters: numOfCards int - the number of cards to deal.
**/
func (re *Deck) Deal(numOfCards int) []Card {
	cardsToDeal := []Card{}

	for i := 0; i < numOfCards; i++ {
		cardsToDeal = append(cardsToDeal, re.Cards[0])
		re.Cards = re.RemoveTopCard()
	}

	return cardsToDeal
}

/**
Name: RemoveTopCard
Purpose: This is a function of the deck to remove the top card from the deck.
This function is used when dealing, drawing, whenever we need to decrease the deck size.
The cards are removed one at a time.
Parameters: None.
**/
func (re *Deck) RemoveTopCard() []Card {
	return append(re.Cards[:0], re.Cards[1:]...)
}

/**
Name: InitializeDeck
Purpose: This function is for creating the specified type of deck.
ex. given "Sabacc" the function will create a deck for the Sabacc game
and return it.
Parameters: deckType string - type of card deck to create and return.
**/
func InitializeDeck(deckType string) Deck {
	//initialize the deck of cards based on game type selected
	deck := Deck{}

	switch deckType {
	case "Pazaak":
		fmt.Println("\nInitializing Pazaak deck...")
		//time.Sleep(3 * time.Second)
		fmt.Println("\nSorry, Pazaak is unavailable at this time.")
		//makePazaakDeck(deck)
	case "Sabacc":
		fmt.Println("\nInitializing Sabacc deck...")
		//time.Sleep(1 * time.Second)
		deck = makeSabaccDeck(deck)
	default:
		fmt.Println("Invalid option. game does not require a deck.")
	}

	return deck
}

/**
Name: ShuffleDeck
Purpose: This function is for shuffling the provided deck of cards.
It will return a new deck of the same cards but shuffled randomly.
Parameters: deck Deck - deck of cards to be shuffled and returned.
**/
func ShuffleDeck(deck Deck) Deck {
	fmt.Println("\nShuffling deck...")
	//time.Sleep(1 * time.Second)
	cards := deck.Cards
	t := time.Now()
	rand.Seed(int64(t.Nanosecond()))
	for i := range cards {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}

	return deck
}

/**
Name: makeSabaccDeck
Purpose: This function is to be used by the InitializeDeck func.
When that Function is provided the deck type of "Sabacc" it will use this
function to make a deck of correllian spike sabacc cards. This deck should
consist of 62 cards. 3 suits (Staves - Circle, Square, Triangle) ranging from
-10 to +10 exluding zero. In addition to the two Slyop cards (value 0).
Parameters: deck Deck - deck of cards to be populated for Sabacc.
**/
func makeSabaccDeck(deck Deck) Deck {

	staves := []string{"circle", "square", "triangle"}
	values := []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, stave := range staves {
		for _, value := range values {
			deck.Cards = append(deck.Cards, Card{Stave: stave, Value: value})
		}
	}
	//add the two sylops (zero cards)
	deck.Cards = append(deck.Cards, Card{Stave: "Sylop", Value: 0})
	deck.Cards = append(deck.Cards, Card{Stave: "Sylop", Value: 0})
	fmt.Printf("\nSabacc deck initialized with %v cards.\n", len(deck.Cards))

	return deck
}

/*
func makePazaakDeck(deck Deck) {

}*/

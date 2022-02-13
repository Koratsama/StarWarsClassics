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

func (re *Deck) Deal(numOfCards int) []Card {
	cardsToDeal := []Card{}

	for i := 0; i < numOfCards; i++ {
		cardsToDeal = append(cardsToDeal, re.Cards[0])
		re.Cards = re.RemoveTopCard()
	}

	return cardsToDeal
}

func (re *Deck) RemoveTopCard() []Card {
	return append(re.Cards[:0], re.Cards[1:]...)
}

func InitializeDeck(deckType string) Deck {
	//initialize the deck of cards based on game type selected
	deck := Deck{}

	switch deckType {
	case "Pazaak":
		fmt.Println("\nInitializing Pazaak deck...")
		time.Sleep(3 * time.Second)
		fmt.Println("\nSorry, Pazaak is unavailable at this time.")
		//makePazaakDeck(deck)
	case "Sabacc":
		fmt.Println("\nInitializing Sabacc deck...")
		time.Sleep(1 * time.Second)
		deck = makeSabaccDeck(deck)
	default:
		fmt.Println("Invalid option. game does not require a deck.")
	}

	return deck
}

func ShuffleDeck(deck Deck) Deck {
	fmt.Println("\nShuffling deck...")
	time.Sleep(1 * time.Second)
	cards := deck.Cards
	t := time.Now()
	rand.Seed(int64(t.Nanosecond()))
	for i := range cards {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}

	return deck
}

/*
func Shuffle(a []int) {
	t := time.Now()
	rand.Seed(int64(t.Nanosecond()))
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}*/

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

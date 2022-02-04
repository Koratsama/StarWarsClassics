package deck

import (
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
}

type Card struct {
	Suit  string
	value int
}

func initializeDeck(deckType string) Deck {
	//initialize the deck of cards based on
	//game type selected
	deck := Deck{}

	return deck
}

func shuffleDeck(deck Deck) Deck {
	cards := deck.Cards
	for i := range cards {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}

	return deck
}

func Shuffle(a []int) {
	t := time.Now()
	rand.Seed(int64(t.Nanosecond()))
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

package player

import "github.com/Koratsama/StarWarsClassics/deck"

type Player struct {
	Name     string
	Hand     []deck.Card
	Credits  int
	Position int
}

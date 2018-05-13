package main

import (
	"fmt"
	"github.com/dancing-koala/gophercises-impl/ex09_card_deck/pkg/deck"
)

func main() {

	cfg := deck.NewConfig()

	cfg.Shuffle = false
	cfg.JokerCount = 3
	cfg.ExtraDecks = 2
	cfg.Filter = func(f, n uint8) bool {
		return f >= 3
	}

	cfg.Sort = func(left, right deck.Card) bool {
		if left.Family > right.Family {
			return true
		}

		return left.Number <= right.Number
	}

	cardDeck := deck.New(cfg)

	fmt.Println(cardDeck)
}

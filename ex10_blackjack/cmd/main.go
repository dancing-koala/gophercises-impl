package main

import (
	"fmt"
	"github.com/dancing-koala/gophercises-impl/ex09_card_deck/pkg/deck"
	"strings"
)

type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))

	for i := range h {
		strs[i] = h[i].String()
	}

	return strings.Join(strs, ", ")
}

func (h Hand) DealerString() string {
	return h[0].String() + ", **HIDDEN**"
}

func (h Hand) Score() int {
	minScore := h.MinScore()

	if minScore > 11 {
		return minScore
	}

	for j := range h {
		if h[j].Number == 1 {
			return minScore + 10
		}
	}

	return minScore
}

func (h Hand) MinScore() int {

	score := 0

	for j := range h {
		score += min(int(h[j].Number), 10)
	}

	return score

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	cfg := deck.NewConfig()
	cfg.ExtraDecks = 2
	cfg.Shuffle = true

	cards := deck.New(cfg)

	var player, dealer Hand
	var card deck.Card

	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&player, &dealer} {
			card, cards = draw(cards)
			*hand = append(*hand, card)
		}
	}

	var input string

	for input != "s" {
		fmt.Println("player: ", player)
		fmt.Println("dealer: ", dealer.DealerString())
		fmt.Println("What will you do? (h)it, (s)tand")

		fmt.Scanf("%s\n", &input)

		switch input {
		case "h":
			card, cards = draw(cards)
			player = append(player, card)
		}
	}

	for dealer.Score() <= 16 || (dealer.Score() == 17 && dealer.MinScore() != 17) {
		card, cards = draw(cards)
		dealer = append(dealer, card)
	}

	pScore, dScore := player.Score(), dealer.Score()

	fmt.Println("==FINAL HANDS==")
	fmt.Println("player: ", player, "Score:", pScore)
	fmt.Println("dealer: ", dealer, "Score:", dScore)

	switch {
	case pScore > 21:
		fmt.Println("Player busted")

	case dScore > 21:
		fmt.Println("Dealer busted")

	case pScore > dScore:
		fmt.Println("Player won")

	case dScore > pScore:
		fmt.Println("Dealer won")

	default:
		fmt.Println("It's a draw!")
	}

}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

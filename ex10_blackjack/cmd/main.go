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

func Shuffle(gs GameState) GameState {
	cfg := deck.NewConfig()
	cfg.ExtraDecks = 2
	cfg.Shuffle = true

	dolly := clone(gs)
	dolly.Deck = deck.New(cfg)

	return dolly
}

func Deal(gs GameState) GameState {
	dolly := clone(gs)

	dolly.Player = make(Hand, 0, 5)
	dolly.Dealer = make(Hand, 0, 5)

	var card deck.Card

	for i := 0; i < 2; i++ {
		card, dolly.Deck = draw(dolly.Deck)
		dolly.Player = append(dolly.Player, card)
		card, dolly.Deck = draw(dolly.Deck)
		dolly.Dealer = append(dolly.Dealer, card)
	}

	return dolly
}

func Stand(gs GameState) GameState {
	dolly := clone(gs)
	dolly.State++

	return dolly
}

func Hit(gs GameState) GameState {
	dolly := clone(gs)

	hand := dolly.CurrentPlayer()

	var card deck.Card
	card, dolly.Deck = draw(dolly.Deck)
	*hand = append(*hand, card)

	if hand.Score() > 21 {
		return Stand(dolly)
	}

	return dolly
}

func EndHand(gs GameState) GameState {
	dolly := clone(gs)

	pScore, dScore := dolly.Player.Score(), dolly.Dealer.Score()

	fmt.Println("==FINAL HANDS==")
	fmt.Println("player: ", dolly.Player, "Score:", pScore)
	fmt.Println("dealer: ", dolly.Dealer, "Score:", dScore)

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

	dolly.Player = nil
	dolly.Dealer = nil

	return dolly
}

func main() {
	var gs GameState

	gs = Shuffle(gs)
	gs = Deal(gs)

	var input string

	for gs.State == StatePlayerTurn {
		fmt.Println("player: ", gs.Player)
		fmt.Println("dealer: ", gs.Dealer.DealerString())
		fmt.Println("What will you do? (h)it, (s)tand")

		fmt.Scanf("%s\n", &input)

		switch input {
		case "h":
			gs = Hit(gs)

		case "s":
			gs = Stand(gs)

		default:
			fmt.Println("Invalid option:", input)
		}
	}

	for gs.State == StateDealerTurn {
		if gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.MinScore() != 17) {
			gs = Hit(gs)
		} else {
			gs = Stand(gs)
		}
	}

	gs = EndHand(gs)
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

type State int8

const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

type GameState struct {
	Deck   []deck.Card
	State  State
	Player Hand
	Dealer Hand
}

func (gs *GameState) CurrentPlayer() *Hand {
	switch gs.State {

	case StatePlayerTurn:
		return &gs.Player

	case StateDealerTurn:
		return &gs.Dealer

	default:
		panic("Not a player's turn !!")
	}
}

func clone(gs GameState) GameState {
	dolly := GameState{
		Deck:   make([]deck.Card, len(gs.Deck)),
		State:  gs.State,
		Player: make(Hand, len(gs.Player)),
		Dealer: make(Hand, len(gs.Dealer)),
	}

	copy(dolly.Deck, gs.Deck)
	copy(dolly.Player, gs.Player)
	copy(dolly.Dealer, gs.Dealer)

	return dolly
}

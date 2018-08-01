package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	SPADE uint8 = iota
	DIAMOND
	CLUB
	HEART
	JOKER
)

type Card struct {
	Family uint8
	Number uint8
}

func (c *Card) Symbol() string {
	switch c.Number {
	case 0:
		return "Jkr"
	case 1:
		return "A"
	case 11:
		return "J"
	case 12:
		return "Q"
	case 13:
		return "K"
	default:
		return fmt.Sprintf("%d", c.Number)
	}
}

type Config struct {
	Sort       func(left, right Card) bool
	Shuffle    bool
	JokerCount uint8
	Filter     func(f, n uint8) bool
	ExtraDecks uint8
}

func NewConfig() *Config {
	return &Config{
		Shuffle:    false,
		JokerCount: 0,
		Filter:     defaultFilter,
		ExtraDecks: 0,
	}
}

func New(cfg *Config) []Card {

	families := []uint8{SPADE, DIAMOND, CLUB, HEART}

	decksCount := 1 + cfg.ExtraDecks

	deck := make([]Card, 0, 52*decksCount+cfg.JokerCount)

	for count := 0; count < int(decksCount); count++ {
		for i := 0; i < 13; i++ {
			for j := 0; j < 4; j++ {

				if !cfg.Filter(families[j], uint8(i+1)) {
					continue
				}

				deck = append(deck, Card{
					Family: families[j],
					Number: uint8(i + 1),
				})
			}
		}
	}

	if cfg.JokerCount > 0 {
		for i := 0; i < int(cfg.JokerCount); i++ {
			deck = append(deck, Card{
				Family: JOKER,
				Number: 0,
			})
		}
	}

	if cfg.Sort != nil {
		sort.Slice(deck, func(i, j int) bool {
			return cfg.Sort(deck[i], deck[j])
		})
	}

	if cfg.Shuffle {
		shuffle(deck)
	}

	return deck
}

func defaultFilter(f, n uint8) bool {
	return true
}

func shuffle(deck []Card) {
	length := len(deck)
	rand.Seed(time.Now().UnixNano())

	for i := length - 1; i >= 0; i-- {
		j := rand.Intn(length)
		deck[i], deck[j] = deck[j], deck[i]
	}
}

func (c Card) String() string {

	var family string

	switch c.Family {
	case SPADE:
		family = "\u2660"
		break

	case DIAMOND:
		family = "\u2666"
		break

	case HEART:
		family = "\u2663"
		break

	case CLUB:
		family = "\u2665"
		break

	}

	return fmt.Sprintf("[%s %s]", c.Symbol(), family)
}

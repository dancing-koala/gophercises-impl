package deck

import (
	"testing"
)

func TestShuffling(t *testing.T) {
	cfg := NewConfig()

	left := New(cfg)

	cfg.Shuffle = true

	right := New(cfg)

	if areSame(left, right) {
		t.Errorf("Decks should be different\nleft=%#v\nright=%#v\n", left, right)
	}
}

func TestJokerCount(t *testing.T) {
	cfg := NewConfig()
	cfg.JokerCount = 3

	deck := New(cfg)

	expectedLen := 52 + int(cfg.JokerCount)

	if len(deck) != expectedLen {
		t.Errorf("Deck should have <%d> cards, got <%d>", expectedLen, len(deck))
	}

	expectedCount := countCardsInFamily(deck, JOKER)

	if expectedCount != int(cfg.JokerCount) {
		t.Errorf("Deck should have <%d> jokers, got <%d>", expectedCount, cfg.JokerCount)
	}
}

func TestExtraDecks(t *testing.T) {
	cfg := NewConfig()

	base := New(cfg)

	cfg.ExtraDecks = 2

	withExtra := New(cfg)

	families := []uint8{SPADE, DIAMOND, CLUB, HEART}

	for _, family := range families {
		baseCount := countCardsInFamily(base, family)
		withExtraCount := countCardsInFamily(withExtra, family)

		expected := int(cfg.ExtraDecks+1) * baseCount

		if withExtraCount != expected {
			t.Errorf("Deck should have <%d> cards of type %d, got <%d>", expected, family, withExtraCount)
		}
	}
}

func TestFilter(t *testing.T) {

	number := uint8(2)

	filterFunc := func(f, n uint8) bool {
		return n != number
	}

	validateFunc := func(f, n uint8) bool {
		return n == number
	}

	cfg := NewConfig()

	base := New(cfg)

	cfg.Filter = filterFunc

	filtered := New(cfg)

	if len(filtered) != 48 {
		t.Errorf("Filtered deck should have 48 cards, got <%d>", len(filtered))
	}

	baseCount := countCardsSatisfying(base, validateFunc)

	if baseCount != 4 {
		t.Errorf("Base deck should have 4 cards with number %d, got <%d>", number, baseCount)
	}

	filteredCount := countCardsSatisfying(filtered, validateFunc)

	if filteredCount > 0 {
		t.Errorf("Filtered deck should have no card with number %d, got <%d>", number, filteredCount)
	}
}

func areSame(left, right []Card) bool {
	if len(left) != len(right) {
		return false
	}

	for i := 0; i < len(left); i++ {
		if left[i].Number != right[i].Number || left[i].Family != right[i].Family {
			return false
		}
	}

	return true
}

func countCardsInFamily(cards []Card, family uint8) int {
	return countCardsSatisfying(cards, func(f, n uint8) bool { return f == family })
}

func countCardsSatisfying(cards []Card, validateFunc func(f, n uint8) bool) int {
	result := 0

	for i := 0; i < len(cards); i++ {
		if validateFunc(cards[i].Family, cards[i].Number) {
			result++
		}
	}

	return result
}

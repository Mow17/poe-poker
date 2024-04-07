package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectStraightFlushCards_1(t *testing.T) {
	cards := []Card{
		{Suit: "s", Number: 6},
		{Suit: "s", Number: 5},
		{Suit: "h", Number: 7},
		{Suit: "s", Number: 1},
		{Suit: "s", Number: 2},
		{Suit: "s", Number: 4},
		{Suit: "s", Number: 3},
	}
	expected, _ := SelectStraightFlushCards(cards)
	correct := []Card{
		{Suit: "s", Number: 6},
		{Suit: "s", Number: 5},
		{Suit: "s", Number: 4},
		{Suit: "s", Number: 3},
		{Suit: "s", Number: 2},
	}
	assert.Equal(t, correct, expected)
}

func TestSelectStraightFlushCards_2(t *testing.T) {
	cards := []Card{
		{Suit: "s", Number: 13},
		{Suit: "h", Number: 5},
		{Suit: "c", Number: 10},
		{Suit: "s", Number: 10},
		{Suit: "s", Number: 12},
		{Suit: "s", Number: 1},
		{Suit: "s", Number: 11},
	}
	expected, _ := SelectStraightFlushCards(cards)
	correct := []Card{
		{Suit: "s", Number: 1},
		{Suit: "s", Number: 13},
		{Suit: "s", Number: 12},
		{Suit: "s", Number: 11},
		{Suit: "s", Number: 10},
	}
	assert.Equal(t, correct, expected)
}

func TestSelectStraightFlushCards_Error(t *testing.T) {
	cards := []Card{
		{Suit: "s", Number: 13},
		{Suit: "h", Number: 5},
		{Suit: "c", Number: 10},
		{Suit: "s", Number: 10},
		{Suit: "s", Number: 12},
		{Suit: "s", Number: 1},
		{Suit: "c", Number: 11},
	}
	_, error := SelectStraightFlushCards(cards)
	assert.NotNil(t, error)
}

func TestSelectFourOfAKindCards_1(t *testing.T) {
	cards := []Card{
		{Suit: "s", Number: 11},
		{Suit: "d", Number: 9},
		{Suit: "c", Number: 11},
		{Suit: "h", Number: 1},
		{Suit: "h", Number: 11},
		{Suit: "h", Number: 7},
		{Suit: "d", Number: 11},
	}
	expected, _ := SelectFourOfAKindCards(cards)
	correct := []Card{
		{Suit: "s", Number: 11},
		{Suit: "c", Number: 11},
		{Suit: "h", Number: 11},
		{Suit: "d", Number: 11},
		{Suit: "h", Number: 1},
	}
	assert.Equal(t, correct, expected)
}

func TestSelectFullHouseCards_1(t *testing.T) {
	cards := []Card{
		{Suit: "s", Number: 11},
		{Suit: "d", Number: 9},
		{Suit: "c", Number: 11},
		{Suit: "h", Number: 1},
		{Suit: "h", Number: 11},
		{Suit: "h", Number: 7},
		{Suit: "d", Number: 7},
	}
	expected, _ := SelectFullHouseCards(cards)
	correct := []Card{
		{Suit: "s", Number: 11},
		{Suit: "c", Number: 11},
		{Suit: "h", Number: 11},
		{Suit: "h", Number: 7},
		{Suit: "d", Number: 7},
	}
	assert.Equal(t, correct, expected)
}

func TestSelectFlushCards_1(t *testing.T) {
	cards := []Card{
		{Suit: "s", Number: 11},
		{Suit: "s", Number: 9},
		{Suit: "s", Number: 1},
		{Suit: "s", Number: 13},
		{Suit: "s", Number: 4},
		{Suit: "s", Number: 2},
		{Suit: "s", Number: 7},
	}
	expected, _ := SelectFlushCards(cards)
	correct := []Card{
		{Suit: "s", Number: 1},
		{Suit: "s", Number: 13},
		{Suit: "s", Number: 11},
		{Suit: "s", Number: 9},
		{Suit: "s", Number: 7},
	}
	assert.Equal(t, correct, expected)
}

func TestSelectStraightCards_1(t *testing.T) {
	cards := []Card{
		{Suit: "s", Number: 11},
		{Suit: "d", Number: 9},
		{Suit: "c", Number: 11},
		{Suit: "h", Number: 1},
		{Suit: "h", Number: 12},
		{Suit: "h", Number: 13},
		{Suit: "h", Number: 10},
	}
	expected, _ := SelectStraightCards(cards)
	correct := []Card{
		{Suit: "h", Number: 1},
		{Suit: "h", Number: 13},
		{Suit: "h", Number: 12},
		{Suit: "s", Number: 11},
		{Suit: "h", Number: 10},
	}
	assert.Equal(t, correct, expected)
}

func TestSelectThreeOfAKindCards_1(t *testing.T) {
	cards := []Card{
		{Suit: "s", Number: 11},
		{Suit: "d", Number: 9},
		{Suit: "c", Number: 11},
		{Suit: "h", Number: 1},
		{Suit: "h", Number: 11},
		{Suit: "h", Number: 7},
		{Suit: "h", Number: 10},
	}
	expected, _ := SelectThreeOfAKindCards(cards)
	correct := []Card{
		{Suit: "s", Number: 11},
		{Suit: "c", Number: 11},
		{Suit: "h", Number: 11},
		{Suit: "h", Number: 1},
		{Suit: "h", Number: 10},
	}
	assert.Equal(t, correct, expected)
}

func TestSelectTwoPairsCards_1(t *testing.T) {
	cards := []Card{
		{Suit: "d", Number: 3},
		{Suit: "h", Number: 12},
		{Suit: "s", Number: 5},
		{Suit: "s", Number: 13},
		{Suit: "s", Number: 9},
		{Suit: "s", Number: 12},
		{Suit: "c", Number: 13},
	}
	expected, _ := SelectTwoPairsCards(cards)
	correct := []Card{
		{Suit: "s", Number: 13},
		{Suit: "c", Number: 13},
		{Suit: "h", Number: 12},
		{Suit: "s", Number: 12},
		{Suit: "s", Number: 9},
	}
	assert.Equal(t, correct, expected)
}

func TestSelectOnePairCards_1(t *testing.T) {
	cards := []Card{
		{Suit: "c", Number: 5},
		{Suit: "s", Number: 10},
		{Suit: "d", Number: 1},
		{Suit: "s", Number: 3},
		{Suit: "c", Number: 4},
		{Suit: "h", Number: 11},
		{Suit: "s", Number: 1},
	}
	expected, _ := SelectOnePairCards(cards)
	correct := []Card{
		{Suit: "d", Number: 1},
		{Suit: "s", Number: 1},
		{Suit: "h", Number: 11},
		{Suit: "s", Number: 10},
		{Suit: "c", Number: 5},
	}
	assert.Equal(t, correct, expected)
}

func TestSelectHighCardCards_1(t *testing.T) {
	cards := []Card{
		{Suit: "s", Number: 5},
		{Suit: "s", Number: 2},
		{Suit: "h", Number: 1},
		{Suit: "h", Number: 10},
		{Suit: "d", Number: 4},
		{Suit: "d", Number: 13},
		{Suit: "c", Number: 6},
	}
	expected, _ := SelectHighCardCards(cards)
	correct := []Card{
		{Suit: "h", Number: 1},
		{Suit: "d", Number: 13},
		{Suit: "h", Number: 10},
		{Suit: "c", Number: 6},
		{Suit: "s", Number: 5},
	}
	assert.Equal(t, correct, expected)
}

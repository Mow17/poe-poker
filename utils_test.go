package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceAceWithOneFromFourteen(t *testing.T) {
	cards := []Card{
		{Suit: "s", Number: 14},
		{Suit: "s", Number: 10},
		{Suit: "s", Number: 11},
		{Suit: "s", Number: 12},
		{Suit: "s", Number: 13},
		{Suit: "h", Number: 14},
		{Suit: "c", Number: 14},
	}
	expected := []Card{
		{Suit: "s", Number: 1},
		{Suit: "s", Number: 10},
		{Suit: "s", Number: 11},
		{Suit: "s", Number: 12},
		{Suit: "s", Number: 13},
		{Suit: "h", Number: 1},
		{Suit: "c", Number: 1},
	}
	actual := ReplaceAceWithOneFromFourteen(cards)
	assert.Equal(t, expected, actual)
}

func TestReplaceAceWithFouteenFromOne(t *testing.T) {
	cards := []Card{
		{Suit: "s", Number: 1},
		{Suit: "s", Number: 10},
		{Suit: "s", Number: 11},
		{Suit: "s", Number: 12},
		{Suit: "s", Number: 13},
		{Suit: "h", Number: 1},
		{Suit: "c", Number: 1},
	}
	expected := []Card{
		{Suit: "s", Number: 14},
		{Suit: "s", Number: 10},
		{Suit: "s", Number: 11},
		{Suit: "s", Number: 12},
		{Suit: "s", Number: 13},
		{Suit: "h", Number: 14},
		{Suit: "c", Number: 14},
	}
	actual := ReplaceAceWithFouteenFromOne(cards)
	assert.Equal(t, expected, actual)
}

func TestAddFourteen(t *testing.T) {
	cards := []Card{
		{Suit: "s", Number: 1},
		{Suit: "s", Number: 2},
		{Suit: "s", Number: 3},
		{Suit: "s", Number: 4},
		{Suit: "s", Number: 5},
		{Suit: "h", Number: 1},
		{Suit: "c", Number: 1},
	}
	expected := []Card{
		{Suit: "s", Number: 1},
		{Suit: "s", Number: 2},
		{Suit: "s", Number: 3},
		{Suit: "s", Number: 4},
		{Suit: "s", Number: 5},
		{Suit: "h", Number: 1},
		{Suit: "c", Number: 1},
		{Suit: "s", Number: 14},
		{Suit: "h", Number: 14},
		{Suit: "c", Number: 14},
	}
	actual := AddFourteen(cards)
	assert.Equal(t, expected, actual)
}

func TestSortIncreasingSuitPriority(t *testing.T) {
	cards := []Card{
		{Suit: "s", Number: 1},
		{Suit: "h", Number: 1},
		{Suit: "c", Number: 1},
		{Suit: "s", Number: 2},
		{Suit: "h", Number: 2},
		{Suit: "c", Number: 2},
		{Suit: "s", Number: 3},
		{Suit: "h", Number: 3},
		{Suit: "c", Number: 3},
		{Suit: "s", Number: 13},
		{Suit: "h", Number: 13},
		{Suit: "c", Number: 13},
		{Suit: "s", Number: 14},
		{Suit: "h", Number: 14},
		{Suit: "c", Number: 14},
	}
	expected := []Card{
		{Suit: "c", Number: 1},
		{Suit: "c", Number: 2},
		{Suit: "c", Number: 3},
		{Suit: "c", Number: 13},
		{Suit: "c", Number: 14},
		{Suit: "h", Number: 1},
		{Suit: "h", Number: 2},
		{Suit: "h", Number: 3},
		{Suit: "h", Number: 13},
		{Suit: "h", Number: 14},
		{Suit: "s", Number: 1},
		{Suit: "s", Number: 2},
		{Suit: "s", Number: 3},
		{Suit: "s", Number: 13},
		{Suit: "s", Number: 14},
	}
	actual := SortIncreasingSuitPriority(cards)
	assert.Equal(t, expected, actual)
}

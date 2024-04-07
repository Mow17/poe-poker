package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortIncreasing_1(t *testing.T) {
	cards := []Card{
		{Suit: "s", Number: 5},
		{Suit: "s", Number: 2},
		{Suit: "s", Number: 1},
		{Suit: "s", Number: 10},
		{Suit: "s", Number: 4},
		{Suit: "s", Number: 1},
		{Suit: "s", Number: 6},
	}
	expected := []Card{
		{Suit: "s", Number: 1},
		{Suit: "s", Number: 1},
		{Suit: "s", Number: 2},
		{Suit: "s", Number: 4},
		{Suit: "s", Number: 5},
		{Suit: "s", Number: 6},
		{Suit: "s", Number: 10},
	}
	actual := SortIncreasing(cards)
	assert.Equal(t, expected, actual)
}

func TestSortIncreasing_2(t *testing.T) {
	cards := []Card{
		{Suit: "s", Number: 13},
		{Suit: "h", Number: 3},
		{Suit: "s", Number: 8},
		{Suit: "s", Number: 2},
		{Suit: "s", Number: 3},
		{Suit: "d", Number: 7},
		{Suit: "s", Number: 1},
	}
	expected := []Card{
		{Suit: "s", Number: 1},
		{Suit: "s", Number: 2},
		{Suit: "h", Number: 3},
		{Suit: "s", Number: 3},
		{Suit: "d", Number: 7},
		{Suit: "s", Number: 8},
		{Suit: "s", Number: 13},
	}
	actual := SortIncreasing(cards)
	assert.Equal(t, expected, actual)
}

func TestSortDecreasing_1(t *testing.T) {
	cards := []Card{
		{Suit: "s", Number: 5},
		{Suit: "s", Number: 2},
		{Suit: "d", Number: 1},
		{Suit: "s", Number: 10},
		{Suit: "s", Number: 4},
		{Suit: "s", Number: 1},
		{Suit: "s", Number: 6},
	}
	expected := []Card{
		{Suit: "s", Number: 10},
		{Suit: "s", Number: 6},
		{Suit: "s", Number: 5},
		{Suit: "s", Number: 4},
		{Suit: "s", Number: 2},
		{Suit: "d", Number: 1},
		{Suit: "s", Number: 1},
	}
	actual := SortDecreasing(cards)
	assert.Equal(t, expected, actual)
}

func TestSortDecreasing_2(t *testing.T) {
	cards := []Card{
		{Suit: "s", Number: 13},
		{Suit: "h", Number: 3},
		{Suit: "s", Number: 8},
		{Suit: "s", Number: 2},
		{Suit: "s", Number: 3},
		{Suit: "d", Number: 7},
		{Suit: "s", Number: 1},
	}
	expected := []Card{
		{Suit: "s", Number: 13},
		{Suit: "s", Number: 8},
		{Suit: "d", Number: 7},
		{Suit: "h", Number: 3},
		{Suit: "s", Number: 3},
		{Suit: "s", Number: 2},
		{Suit: "s", Number: 1},
	}
	actual := SortDecreasing(cards)
	assert.Equal(t, expected, actual)
}

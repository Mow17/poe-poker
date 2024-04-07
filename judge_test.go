package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJudgeHandRank1(t *testing.T) {
	cards := []Card{
		{Suit: "d", Number: 5},
		{Suit: "s", Number: 2},
		{Suit: "h", Number: 1},
		{Suit: "s", Number: 10},
		{Suit: "s", Number: 4},
		{Suit: "c", Number: 1},
		{Suit: "h", Number: 6},
	}
	rank := JudgeHandRank(cards)
	assert.Equal(t, OnePair, rank)
}

func TestJudgeHandRank2(t *testing.T) {
	cards := []Card{
		{Suit: "d", Number: 13},
		{Suit: "d", Number: 3},
		{Suit: "c", Number: 8},
		{Suit: "s", Number: 2},
		{Suit: "s", Number: 3},
		{Suit: "s", Number: 7},
		{Suit: "d", Number: 1},
	}
	rank := JudgeHandRank(cards)
	assert.Equal(t, OnePair, rank)
}

func TestJudgeHandRank3(t *testing.T) {
	cards := []Card{
		{Suit: "s", Number: 11},
		{Suit: "d", Number: 7},
		{Suit: "h", Number: 5},
		{Suit: "c", Number: 1},
		{Suit: "s", Number: 13},
		{Suit: "s", Number: 4},
		{Suit: "d", Number: 2},
	}
	cards = SortIncreasing(cards)
	rank := JudgeHandRank(cards)
	assert.Equal(t, HighCard, rank)
}

func TestJudgeStraight(t *testing.T) {
	cards := []Card{
		{Suit: "s", Number: 11},
		{Suit: "d", Number: 9},
		{Suit: "c", Number: 11},
		{Suit: "h", Number: 1},
		{Suit: "h", Number: 11},
		{Suit: "h", Number: 7},
		{Suit: "h", Number: 10},
	}
	flag := JudgeStraight(cards)
	assert.Equal(t, false, flag)
}

func TestJudgeThreeOfAKind(t *testing.T) {
	cards := []Card{
		{Suit: "s", Number: 11},
		{Suit: "d", Number: 9},
		{Suit: "c", Number: 11},
		{Suit: "h", Number: 1},
		{Suit: "h", Number: 11},
		{Suit: "h", Number: 7},
		{Suit: "h", Number: 10},
	}
	flag := JudgeThreeOfAKind(cards)
	assert.Equal(t, true, flag)
}

func TestJudgeTwoPairs(t *testing.T) {
	cards := []Card{
		{Suit: "h", Number: 12},
		{Suit: "d", Number: 12},
		{Suit: "d", Number: 1},
		{Suit: "h", Number: 6},
		{Suit: "d", Number: 5},
		{Suit: "s", Number: 5},
		{Suit: "s", Number: 10},
	}
	flag := JudgeTwoPairs(cards)
	assert.Equal(t, true, flag)
}

func TestJudgeOnePair1(t *testing.T) {
	cards := []Card{
		{Suit: "c", Number: 5},
		{Suit: "s", Number: 10},
		{Suit: "d", Number: 1},
		{Suit: "s", Number: 3},
		{Suit: "c", Number: 4},
		{Suit: "h", Number: 11},
		{Suit: "s", Number: 1},
	}
	flag := JudgeOnePair(cards)
	assert.Equal(t, true, flag)
}

func TestJudgeOnePair2(t *testing.T) {
	cards := []Card{
		{Suit: "d", Number: 5},
		{Suit: "s", Number: 2},
		{Suit: "h", Number: 1},
		{Suit: "s", Number: 10},
		{Suit: "s", Number: 4},
		{Suit: "c", Number: 1},
		{Suit: "h", Number: 6},
	}
	flag := JudgeOnePair(cards)
	assert.Equal(t, true, flag)
}

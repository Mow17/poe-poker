package main

import "sort"

func ReplaceAceWithOneFromFourteen(cards []Card) []Card {
	for i, card := range cards {
		if card.Number == 14 {
			cards[i].Number = 1
		}
	}
	return cards
}

func ReplaceAceWithFouteenFromOne(cards []Card) []Card {
	for i, card := range cards {
		if card.Number == 1 {
			cards[i].Number = 14
		}
	}
	return cards
}

func AddFourteen(cards []Card) []Card {
	for _, card := range cards {
		if card.Number == 1 {
			cards = append(cards, Card{Suit: card.Suit, Number: 14})
		}
	}
	return cards
}

func ConvertRankNumberToRank(rankNumber HandRank) string {
	rank := ""
	switch rankNumber {
	case 0:
		rank = "HighCard"
	case 1:
		rank = "OnePair"
	case 2:
		rank = "TwoPairs"
	case 3:
		rank = "ThreeOfAKind"
	case 4:
		rank = "Straight"
	case 5:
		rank = "Flush"
	case 6:
		rank = "FullHouse"
	case 7:
		rank = "FourOfAKind"
	case 8:
		rank = "StraightFlush"
	}
	return rank
}

func SortIncreasingSuitPriority(cards []Card) []Card {
	sort.SliceStable(cards, func(i, j int) bool {
		// If suits are different, compare by suit priority
		if cards[i].Suit != cards[j].Suit {
			return cards[i].Suit < cards[j].Suit
		}
		// If the suits are the same, compare by number
		return cards[i].Number < cards[j].Number
	})
	return cards
}

func SortDecreaseingSuitPriority(cards []Card) []Card {
	sort.SliceStable(cards, func(i, j int) bool {
		// If suits are different, compare by suit priority
		if cards[i].Suit != cards[j].Suit {
			return cards[i].Suit > cards[j].Suit
		}
		// If the suits are the same, compare by number
		return cards[i].Number > cards[j].Number
	})
	return cards
}

func SortDecreasing(cards []Card) []Card {
	sortedCards := make([]Card, len(cards))
	copy(sortedCards, cards)
	sort.SliceStable(sortedCards, func(i, j int) bool {
		return sortedCards[i].Number > sortedCards[j].Number
	})
	return sortedCards
}

func SortIncreasing(cards []Card) []Card {
	sortedCards := make([]Card, len(cards))
	copy(sortedCards, cards)
	sort.SliceStable(sortedCards, func(i, j int) bool {
		return sortedCards[i].Number < sortedCards[j].Number
	})
	return sortedCards
}

func IsCardsValid(cards []Card, length int) bool {
	if len(cards) != length {
		return false
	}
	// Check if the cards are unique
	uniqueCards := make(map[Card]bool)
	for _, card := range cards {
		if card.Number < 1 || card.Number > 13 {
			return false
		}
		if _, ok := uniqueCards[card]; ok {
			return false
		}
		uniqueCards[card] = true
	}

	return true
}

func AreHandEqual(cards1 []Card, cards2 []Card) bool {
	if len(cards1) != len(cards2) {
		return false
	}
	cards1 = SortIncreasing(cards1)
	cards2 = SortIncreasing(cards2)
	for i := 0; i < len(cards1); i++ {
		if cards1[i].Number != cards2[i].Number {
			return false
		}
	}
	return true
}

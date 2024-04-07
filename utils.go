package main

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

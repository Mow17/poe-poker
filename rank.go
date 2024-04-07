package main

import (
	"errors"
	"sort"
)

func SortBySuitPriority(cards []Card) []Card {
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

func JudgeStraightFlush(cards []Card) bool {
	for _, card := range cards {
		if card.Number == 1 {
			cards = append(cards, Card{Suit: card.Suit, Number: 14})
		}
	}
	cards = SortBySuitPriority(cards)
	straightFlush := false
	previousSuit := cards[0].Suit
	serialCount := 1
	previousNumber := cards[0].Number

	for i, card := range cards {
		if i == 0 {
			continue
		}
		if card.Suit == previousSuit && card.Number == previousNumber+1 {
			serialCount++
		} else {
			serialCount = 1
		}

		if serialCount == 5 {
			straightFlush = true
			break
		}
		previousSuit = card.Suit
		previousNumber = card.Number
	}
	return straightFlush
}

func JudgeFourOfAKind(cards []Card) bool {
	fourOfAKind := false
	numCounts := map[int]int{}
	for _, card := range cards {
		numCounts[card.Number]++
	}

	for _, count := range numCounts {
		if count == 4 {
			fourOfAKind = true
		}
	}
	return fourOfAKind
}

func JudgeFullHouse(cards []Card) bool {
	fullHouse := false
	numCounts := map[int]int{}
	for _, card := range cards {
		numCounts[card.Number]++
	}

	threeOfAKind := false
	twoPairs := false
	for _, count := range numCounts {
		if count == 3 {
			threeOfAKind = true
		} else if count == 2 {
			twoPairs = true
		}
	}

	if threeOfAKind && twoPairs {
		fullHouse = true
	}
	return fullHouse
}

func JudgeFlush(cards []Card) bool {
	flush := false
	suitsCount := map[string]int{}
	for _, card := range cards {
		suitsCount[card.Suit]++
	}

	for _, count := range suitsCount {
		if count >= 5 {
			flush = true
		}
	}
	return flush
}

func JudgeStraight(cards []Card) bool {
	cards = AddFourteen(cards)
	cards = SortDecreasing(cards)
	straight := false
	serialCount := 1
	previousNumber := cards[0].Number

	for i, card := range cards {
		if i == 0 {
			continue
		}
		if card.Number == previousNumber-1 {
			serialCount++
		} else if card.Number == previousNumber {
			continue
		} else {
			serialCount = 1
		}

		if serialCount == 5 {
			straight = true
			break
		}
		previousNumber = card.Number
	}
	return straight
}

func JudgeThreeOfAKind(cards []Card) bool {
	threeOfAKind := false
	numCounts := map[int]int{}
	for _, card := range cards {
		numCounts[card.Number]++
	}

	for _, count := range numCounts {
		if count == 3 {
			threeOfAKind = true
		}
	}
	return threeOfAKind
}

func JudgeTwoPairs(cards []Card) bool {
	twoPairs := false
	numCounts := map[int]int{}
	for _, card := range cards {
		numCounts[card.Number]++
	}

	pairCount := 0
	for _, count := range numCounts {
		if count == 2 {
			pairCount++
		}
	}

	if pairCount == 2 {
		twoPairs = true
	}
	return twoPairs
}

func JudgeOnePair(cards []Card) bool {
	onePair := false
	numCounts := map[int]int{}
	for _, card := range cards {
		numCounts[card.Number]++
	}

	for _, count := range numCounts {
		if count == 2 {
			onePair = true
		}
	}
	return onePair
}

// Warning: Assuming that there is only one suit in which flush may be completed.
func SelectStraightFlushCards(cards []Card) ([]Card, error) {
	cards = AddFourteen(cards)
	cards = SortDecreaseingSuitPriority(cards)

	hand := []Card{}
	tmp_count := 1
	tmp_card := cards[0]
	for i, card := range cards {
		if i == 0 {
			continue
		}
		if card.Number == tmp_card.Number-1 && card.Suit == tmp_card.Suit {
			tmp_count++
			if tmp_count == 5 {
				hand = append(hand, cards[i-4:i+1]...)
				break
			}
		} else {
			tmp_count = 1
		}
		tmp_card = card
	}
	if len(hand) == 0 {
		return nil, errors.New("error: No straight flush found")
	}

	hand = ReplaceAceWithOneFromFourteen(hand)

	return hand, nil
}

// Warnings: The order of the four-card suit is not guaranteed
func SelectFourOfAKindCards(cards []Card) ([]Card, error) {
	cards = AddFourteen(cards)

	numCounts := map[int]int{}
	for _, card := range cards {
		numCounts[card.Number]++
	}

	// Identify the numbers on the four cards
	fourOfAKindNumber := 0
	for number, count := range numCounts {
		if count == 4 {
			fourOfAKindNumber = number
			break
		}
	}
	if fourOfAKindNumber == 0 {
		return nil, errors.New("error: No four of a kind found")
	}

	hand := []Card{}
	highestCard := Card{Number: 0}
	for _, card := range cards {
		if card.Number == fourOfAKindNumber {
			hand = append(hand, card)
		} else if card.Number > highestCard.Number {
			highestCard = card
		}
	}
	hand = append(hand, highestCard)
	hand = ReplaceAceWithOneFromFourteen(hand)

	return hand, nil
}

// Note: Three-card comes before two-pair
// Warnings: The order of the three-card and two-pair suit are not guaranteed
func SelectFullHouseCards(cards []Card) ([]Card, error) {
	cards = ReplaceAceWithFouteenFromOne(cards)

	numCounts := map[int]int{}
	for _, card := range cards {
		numCounts[card.Number]++
	}

	// Identify highest numbers on three cards and two cards
	tmpHighestNumberOfThreeCards := 0
	tmpHighestNumberOfTwoCards := 0
	for number, count := range numCounts {
		if count == 3 && number > tmpHighestNumberOfThreeCards {
			tmpHighestNumberOfThreeCards = number
		} else if count == 2 && number > tmpHighestNumberOfTwoCards {
			tmpHighestNumberOfTwoCards = number
		}
	}

	if tmpHighestNumberOfThreeCards == 0 || tmpHighestNumberOfTwoCards == 0 {
		return nil, errors.New("error: No full house found")
	}

	// Select the highest three-card and two-card
	hand := []Card{}
	for _, card := range cards {
		if card.Number == tmpHighestNumberOfThreeCards {
			hand = append(hand, card)
		}
	}
	for _, card := range cards {
		if card.Number == tmpHighestNumberOfTwoCards {
			hand = append(hand, card)
		}
	}

	hand = ReplaceAceWithOneFromFourteen(hand)

	return hand, nil
}

func SelectFlushCards(cards []Card) ([]Card, error) {
	cards = AddFourteen(cards)
	cards = SortDecreaseingSuitPriority(cards)

	// Identify the suit with flush
	suitCounts := map[string]int{}
	for _, card := range cards {
		suitCounts[card.Suit]++
	}
	suit := ""
	for s, count := range suitCounts {
		if count >= 5 {
			suit = s
			break
		}
	}

	if suit == "" {
		return nil, errors.New("error: No flush found")
	}

	hand := []Card{}
	count := 0
	for _, card := range cards {
		if card.Suit == suit {
			hand = append(hand, card)
			count++
		}
		if count == 5 {
			break
		}
	}

	hand = ReplaceAceWithOneFromFourteen(hand)

	return hand, nil
}

func SelectStraightCards(cards []Card) ([]Card, error) {
	cards = AddFourteen(cards)
	cards = SortDecreasing(cards)

	hand := []Card{cards[0]}
	for i, card := range cards {
		if i == 0 {
			continue
		}
		if card.Number == hand[len(hand)-1].Number-1 {
			hand = append(hand, card)
			if len(hand) == 5 {
				break
			}
		} else if card.Number != hand[len(hand)-1].Number {
			hand = []Card{card}
		}
	}
	if len(hand) == 0 {
		return nil, errors.New("error: No straight found")
	}

	hand = ReplaceAceWithOneFromFourteen(hand)

	return hand, nil
}

func SelectThreeOfAKindCards(cards []Card) ([]Card, error) {
	cards = ReplaceAceWithFouteenFromOne(cards)

	numCounts := map[int]int{}
	for _, card := range cards {
		numCounts[card.Number]++
	}

	// Identify the number on the three cards
	threeOfAKindNumber := 0
	for number, count := range numCounts {
		if count == 3 {
			threeOfAKindNumber = number
			break
		}
	}
	if threeOfAKindNumber == 0 {
		return nil, errors.New("error: No three of a kind found")
	}

	hand := []Card{}
	highestCards := []Card{}
	for _, card := range cards {
		if card.Number == threeOfAKindNumber {
			hand = append(hand, card)
		} else {
			highestCards = append(highestCards, card)
		}
	}
	sort.Slice(highestCards, func(i, j int) bool {
		return highestCards[i].Number > highestCards[j].Number
	})
	hand = append(hand, highestCards[:2]...)
	hand = ReplaceAceWithOneFromFourteen(hand)

	return hand, nil
}

func SelectTwoPairsCards(cards []Card) ([]Card, error) {
	cards = ReplaceAceWithFouteenFromOne(cards)

	numCounts := map[int]int{}
	for _, card := range cards {
		numCounts[card.Number]++
	}

	// Identify the numbers on the two pairs
	tmpHighestNumberOfFirstPair := 0
	tmpHighestNumberOfSecondPair := 0
	for number, count := range numCounts {
		if count == 2 {
			if number > tmpHighestNumberOfFirstPair {
				tmpHighestNumberOfSecondPair = tmpHighestNumberOfFirstPair
				tmpHighestNumberOfFirstPair = number
			} else if number > tmpHighestNumberOfSecondPair {
				tmpHighestNumberOfSecondPair = number
			}
		}
	}

	if tmpHighestNumberOfFirstPair == 0 || tmpHighestNumberOfSecondPair == 0 {
		return nil, errors.New("error: No two pairs found")
	}

	// Select the highest two pairs
	hand := []Card{}
	highestCard := Card{Number: 0}
	for _, card := range cards {
		if card.Number == tmpHighestNumberOfFirstPair {
			hand = append(hand, card)
		} else if card.Number != tmpHighestNumberOfSecondPair && card.Number > highestCard.Number {
			highestCard = card
		}
	}

	for _, card := range cards {
		if card.Number == tmpHighestNumberOfSecondPair {
			hand = append(hand, card)
		}
	}
	hand = append(hand, highestCard)
	hand = ReplaceAceWithOneFromFourteen(hand)

	return hand, nil
}

func SelectOnePairCards(cards []Card) ([]Card, error) {
	cards = ReplaceAceWithFouteenFromOne(cards)
	numCounts := map[int]int{}
	for _, card := range cards {
		numCounts[card.Number]++
	}

	// Identify the number on the pair
	onePairNumber := 0
	for number, count := range numCounts {
		if count == 2 {
			onePairNumber = number
			break
		}
	}
	if onePairNumber == 0 {
		return nil, errors.New("error: No one pair found")
	}

	hand := []Card{}
	highestCards := []Card{}
	for _, card := range cards {
		if card.Number == onePairNumber {
			hand = append(hand, card)
		} else {
			highestCards = append(highestCards, card)
		}
	}
	sort.Slice(highestCards, func(i, j int) bool {
		return highestCards[i].Number > highestCards[j].Number
	})
	hand = append(hand, highestCards[:3]...)
	hand = ReplaceAceWithOneFromFourteen(hand)

	return hand, nil
}

func SelectHighCardCards(cards []Card) ([]Card, error) {
	cards = AddFourteen(cards)

	// Sort cards in descending order
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Number > cards[j].Number
	})

	hand := ReplaceAceWithOneFromFourteen(cards[:5])

	return hand, nil
}

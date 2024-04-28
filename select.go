package main

import (
	"errors"
	"fmt"
	"log"
)

func SelectStrongestCards(cards []Card, rank HandRank) []Card {
	cards_copy := make([]Card, len(cards))
	copy(cards_copy, cards)
	switch rank {
	case StraightFlush:
		selectedCards, err := SelectStraightFlushCards(cards_copy)
		if err != nil {
			log.Panic(err)
		}
		return selectedCards
	case FourOfAKind:
		selectedCards, err := SelectFourOfAKindCards(cards_copy)
		if err != nil {
			log.Panic(err)
		}
		return selectedCards
	case FullHouse:
		selectedCards, err := SelectFullHouseCards(cards_copy)
		if err != nil {
			log.Panic(err)
		}
		return selectedCards
	case Flush:
		selectedCards, err := SelectFlushCards(cards_copy)
		if err != nil {
			log.Panic(err)
		}
		return selectedCards
	case Straight:
		selectedCards, err := SelectStraightCards(cards_copy)
		if err != nil {
			log.Panic(err)
		}
		return selectedCards
	case ThreeOfAKind:
		selectedCards, err := SelectThreeOfAKindCards(cards_copy)
		if err != nil {
			log.Panic(err)
		}
		return selectedCards
	case TwoPairs:
		selectedCards, err := SelectTwoPairsCards(cards_copy)
		if err != nil {
			log.Panic(err)
		}
		return selectedCards
	case OnePair:
		selectedCards, err := SelectOnePairCards(cards_copy)
		if err != nil {
			log.Panic(err)
		}
		return selectedCards
	default:
		selectedCards, err := SelectHighCardCards(cards_copy)
		if err != nil {
			log.Panic(err)
		}
		return selectedCards
	}
}

// Warning: Assuming that there is only one suit in which flush may be completed.
func SelectStraightFlushCards(cards []Card) ([]Card, error) {
	cards_copy := make([]Card, len(cards))
	copy(cards_copy, cards)

	cards_copy = AddFourteen(cards_copy)
	cards_copy = SortDecreaseingSuitPriority(cards_copy)

	hand := []Card{}
	tmp_count := 1
	tmp_card := cards_copy[0]
	for i, card := range cards_copy {
		if i == 0 {
			continue
		}
		if card.Number == tmp_card.Number-1 && card.Suit == tmp_card.Suit {
			tmp_count++
			if tmp_count == 5 {
				hand = append(hand, cards_copy[i-4:i+1]...)
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
// Warnings: Four-card must be one kind
// Warnings: The length of cards must be 7 or less
func SelectFourOfAKindCards(cards []Card) ([]Card, error) {
	cards_copy := make([]Card, len(cards))
	copy(cards_copy, cards)

	cards_copy = ReplaceAceWithFouteenFromOne(cards_copy)

	numCounts := map[int]int{}
	for _, card := range cards_copy {
		numCounts[card.Number]++
	}

	// Identify the numbers on the four cards_copy
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
	for _, card := range cards_copy {
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
	cards_copy := make([]Card, len(cards))
	copy(cards_copy, cards)

	cards_copy = ReplaceAceWithFouteenFromOne(cards_copy)

	numCounts := map[int]int{}
	for _, card := range cards_copy {
		numCounts[card.Number]++
	}

	// Identify highest numbers on three cards and two cards
	tmpHighestNumberOfThreeCards := 0
	for number, count := range numCounts {
		if count == 3 && number > tmpHighestNumberOfThreeCards {
			tmpHighestNumberOfThreeCards = number
		}
	}
	tmpHighestNumberOfTwoCards := 0
	for number, count := range numCounts {
		if count >= 2 && number > tmpHighestNumberOfTwoCards && number != tmpHighestNumberOfThreeCards {
			tmpHighestNumberOfTwoCards = number
		}
	}

	if tmpHighestNumberOfThreeCards == 0 || tmpHighestNumberOfTwoCards == 0 {
		return nil, errors.New("error: No full house found")
	}

	// Select the highest three-card and two-card
	hand := []Card{}
	for _, card := range cards_copy {
		if card.Number == tmpHighestNumberOfThreeCards {
			hand = append(hand, card)
		}
	}
	countAdded := 0
	for _, card := range cards_copy {
		if card.Number == tmpHighestNumberOfTwoCards {
			hand = append(hand, card)
			countAdded++
		}
		if countAdded == 2 {
			break
		}
	}

	hand = ReplaceAceWithOneFromFourteen(hand)
	return hand, nil
}

// Warnings: Flush must be only one suit
func SelectFlushCards(cards []Card) ([]Card, error) {
	cards_copy := make([]Card, len(cards))
	copy(cards_copy, cards)

	cards_copy = ReplaceAceWithFouteenFromOne(cards_copy)
	cards_copy = SortDecreaseingSuitPriority(cards_copy)

	// Identify the suit with flush
	suitCounts := map[string]int{}
	for _, card := range cards_copy {
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
	for _, card := range cards_copy {
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
	cards_copy := make([]Card, len(cards))
	copy(cards_copy, cards)

	cards_copy = AddFourteen(cards_copy)
	cards_copy = SortDecreasing(cards_copy)

	hand := []Card{cards_copy[0]}
	for i, card := range cards_copy {
		if i == 0 {
			continue
		}
		if card.Number == hand[len(hand)-1].Number-1 {
			hand = append(hand, card)
			if len(hand) == 5 {
				break
			}
		} else if card.Number == hand[len(hand)-1].Number {
			continue
		} else {
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
	cards_copy := make([]Card, len(cards))
	copy(cards_copy, cards)

	cards_copy = ReplaceAceWithFouteenFromOne(cards_copy)
	cards_copy = SortDecreasing(cards_copy)

	numCounts := map[int]int{}
	for _, card := range cards_copy {
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
	for _, card := range cards_copy {
		if card.Number == threeOfAKindNumber {
			hand = append(hand, card)
		}
	}
	countAdded := 0
	for _, card := range cards_copy {
		if card.Number != threeOfAKindNumber {
			hand = append(hand, card)
			countAdded++
		}
		if countAdded == 2 {
			break
		}
	}

	hand = ReplaceAceWithOneFromFourteen(hand)
	for _, card := range hand {
		if card.Number == 14 {
			fmt.Println("error: Ace is not replaced with one; SelectThreeOfAKindCards")
		}
	}
	return hand, nil
}

func SelectTwoPairsCards(cards []Card) ([]Card, error) {
	cards_copy := make([]Card, len(cards))
	copy(cards_copy, cards)

	cards_copy = ReplaceAceWithFouteenFromOne(cards_copy)
	cards_copy = SortDecreasing(cards_copy)

	numCounts := map[int]int{}
	for _, card := range cards_copy {
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
	for _, card := range cards_copy {
		if card.Number == tmpHighestNumberOfFirstPair {
			hand = append(hand, card)
		}
	}
	for _, card := range cards_copy {
		if card.Number == tmpHighestNumberOfSecondPair {
			hand = append(hand, card)
		}
	}
	for _, card := range cards_copy {
		if card.Number != tmpHighestNumberOfFirstPair && card.Number != tmpHighestNumberOfSecondPair {
			hand = append(hand, card)
			break
		}
	}

	hand = ReplaceAceWithOneFromFourteen(hand)
	return hand, nil
}

func SelectOnePairCards(cards []Card) ([]Card, error) {
	cards_copy := make([]Card, len(cards))
	copy(cards_copy, cards)

	cards_copy = ReplaceAceWithFouteenFromOne(cards_copy)
	cards_copy = SortDecreasing(cards_copy)

	numCounts := map[int]int{}
	for _, card := range cards_copy {
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
	for _, card := range cards_copy {
		if card.Number == onePairNumber {
			hand = append(hand, card)
		}
	}
	countAdded := 0
	for _, card := range cards_copy {
		if card.Number != onePairNumber {
			hand = append(hand, card)
			countAdded++
		}
		if countAdded == 3 {
			break
		}
	}
	hand = ReplaceAceWithOneFromFourteen(hand)
	return hand, nil
}

func SelectHighCardCards(cards []Card) ([]Card, error) {
	cards_copy := make([]Card, len(cards))
	copy(cards_copy, cards)

	cards_copy = ReplaceAceWithFouteenFromOne(cards_copy)
	cards_copy = SortDecreasing(cards_copy)

	hand := cards_copy[:5]

	hand = ReplaceAceWithOneFromFourteen(hand)
	cards_copy = ReplaceAceWithOneFromFourteen(cards_copy)
	for _, card := range cards_copy {
		if card.Number == 14 {
			fmt.Println("error: Ace is not replaced with one in cards; SelectHighCardCards")
		}
	}
	for _, card := range hand {
		if card.Number == 14 {
			fmt.Println("error: Ace is not replaced with one in hand; SelectHighCardCards")
		}
	}
	return hand, nil
}

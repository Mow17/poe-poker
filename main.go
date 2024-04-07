package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type HandRank int

const (
	HighCard HandRank = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
)

type Hand struct {
	pid   int
	Rank  HandRank
	Cards []Card
}

func JudgeHandRank(cards []Card) HandRank {
	cards = SortIncreasing(cards)
	handRank := HighCard
	switch {
	case JudgeStraightFlush(cards):
		handRank = StraightFlush
	case JudgeFourOfAKind(cards):
		handRank = FourOfAKind
	case JudgeFullHouse(cards):
		handRank = FullHouse
	case JudgeFlush(cards):
		handRank = Flush
	case JudgeStraight(cards):
		handRank = Straight
	case JudgeThreeOfAKind(cards):
		handRank = ThreeOfAKind
	case JudgeTwoPairs(cards):
		handRank = TwoPairs
	case JudgeOnePair(cards):
		handRank = OnePair
	default:
		handRank = HighCard
	}
	return handRank
}

func SelectStrongestCards(cards []Card, rank HandRank) []Card {
	selectedCards := make([]Card, 5)
	switch rank {
	case StraightFlush:
		selectedCards, _ = SelectStraightFlushCards(cards)
	case FourOfAKind:
		selectedCards, _ = SelectFourOfAKindCards(cards)
	case FullHouse:
		selectedCards, _ = SelectFullHouseCards(cards)
	case Flush:
		selectedCards, _ = SelectFlushCards(cards)
	case Straight:
		selectedCards, _ = SelectStraightCards(cards)
	case ThreeOfAKind:
		selectedCards, _ = SelectThreeOfAKindCards(cards)
	case TwoPairs:
		selectedCards, _ = SelectTwoPairsCards(cards)
	case OnePair:
		selectedCards, _ = SelectOnePairCards(cards)
	case HighCard:
		selectedCards, _ = SelectHighCardCards(cards)
	}
	selectedCards = ReplaceAceWithOneFromFourteen(selectedCards)
	return selectedCards
}

func DetermineHandRank(cards []Card) (Hand, error) {
	if len(cards) < 5 {
		return Hand{}, errors.New("given cards are less than 5 and no rank can be determined")
	}

	// Initialize hand
	hand := Hand{Rank: HighCard, Cards: []Card{}}
	hand.Rank = JudgeHandRank(cards)
	hand.Cards = SelectStrongestCards(cards, hand.Rank)

	return hand, nil
}

func CalculateHandProbability() {
	start := time.Now()
	statisticsHandRank := make(map[HandRank]int)

	// Trials=100000 -> Time taken: 1.5s
	// Trials=1000000 -> Time taken: 23.5s
	// Trials=10000000 -> Time taken: 2m25.3s
	trials := 100000

	for i := 0; i < trials; i++ {
		deck := NewDeck()
		deck.Shuffle()
		cards, _ := deck.DrawCards(7)

		hand, _ := DetermineHandRank(cards)
		statisticsHandRank[hand.Rank]++
	}

	for rank, count := range statisticsHandRank {
		percentage := float64(count) / float64(trials) * 100
		fmt.Printf("%s: %.4f%%\n", ConvertRankNumberToRank(rank), percentage)
	}

	elapsed := time.Since(start)
	fmt.Printf("Time taken: %s\n", elapsed)
	// Print the average time taken for each trial
	fmt.Printf("Average time taken for each trial: %s\n", elapsed/time.Duration(trials))
}

// Warnings: Length of hand1.Cards and hand2.Cards must be 5
// Return 1 if hand1 wins, 2 if hand2 wins, 0 if it's a tie
func CompareTwoHands(hand1 Hand, hand2 Hand) (int, error) {
	if len(hand1.Cards) != 5 || len(hand2.Cards) != 5 {
		return -1, errors.New("length of hand1.Cards and hand2.Cards must be 5")

	}
	// Return 1 if hand1 wins, 2 if hand2 wins, 0 if it's a tie
	if hand1.Rank > hand2.Rank {
		return 1, nil
	} else if hand1.Rank < hand2.Rank {
		return 2, nil
	}

	// Compare the cards of the same rank
	hand1.Cards = ReplaceAceWithFouteenFromOne(hand1.Cards)
	hand2.Cards = ReplaceAceWithFouteenFromOne(hand2.Cards)
	for i := 0; i < 5; i++ {
		if hand1.Cards[i].Number > hand2.Cards[i].Number {
			return 1, nil
		} else if hand1.Cards[i].Number < hand2.Cards[i].Number {
			return 2, nil
		}
	}
	return 0, nil
}

// Warnings: Length of hands[i].Cards must be 5
// O(N^2) time complexity
func DetermineHandWinner(hands []Hand) ([]Hand, error) {
	if len(hands) == 0 {
		return nil, errors.New("no hands to compare")
	}

	// Assign player id to each hand
	for i, hand := range hands {
		hand.pid = i + 1
	}

	// Sort hands in descending order
	// The strongest hand will be at the beginning
	sortedHands := make([]Hand, len(hands))
	copy(sortedHands, hands)
	for i := 0; i < len(sortedHands); i++ {
		for j := i + 1; j < len(sortedHands); j++ {
			resultCompareTwoHands, _ := CompareTwoHands(sortedHands[i], sortedHands[j])
			if resultCompareTwoHands == 2 {
				sortedHands[i], sortedHands[j] = sortedHands[j], sortedHands[i]
			}
		}
	}
	return sortedHands, nil
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

func PrintRandomHandResult() {
	deck := NewDeck()
	deck.Shuffle()
	cards_p1, _ := deck.DrawCards(2)
	cards_p2, _ := deck.DrawCards(2)
	board, _ := deck.DrawCards(5)
	fmt.Println("Player 1's cards:", cards_p1)
	fmt.Println("Player 2's cards:", cards_p2)
	fmt.Println("Board cards:", board)

	hand_p1, _ := DetermineHandRank(append(cards_p1, board...))
	hand_p2, _ := DetermineHandRank(append(cards_p2, board...))

	sortedHands, _ := DetermineHandWinner([]Hand{hand_p1, hand_p2})
	fmt.Println("Player", sortedHands[0].pid, "wins with", ConvertRankNumberToRank(sortedHands[0].Rank), sortedHands[0].Cards)
	fmt.Println("Player", sortedHands[1].pid, "loses with", ConvertRankNumberToRank(sortedHands[1].Rank), sortedHands[1].Cards)
}

func ConvertNumberToSymbol(number int) string {
	switch number {
	case 1:
		return "A"
	case 11:
		return "J"
	case 12:
		return "Q"
	case 13:
		return "K"
	case 14:
		return "A"
	default:
		return strconv.Itoa(number)
	}
}

// the lengthe of cards must be 2
func ConvertCardsToSymbol(cards []Card) (string, error) {
	if len(cards) != 2 {
		return "", errors.New("the length of cards must be 2")
	}
	cards = ReplaceAceWithFouteenFromOne(cards)
	biggerNumber := cards[0].Number
	smallerNumber := cards[1].Number
	if cards[0].Number < cards[1].Number {
		biggerNumber, smallerNumber = smallerNumber, biggerNumber
	}

	biggerNumber_str := ConvertNumberToSymbol(biggerNumber)
	smallerNumber_str := ConvertNumberToSymbol(smallerNumber)
	cards_str := biggerNumber_str + smallerNumber_str
	if cards[0].Suit == cards[1].Suit {
		cards_str += "s"
	} else {
		cards_str += "o"
	}
	return cards_str, nil
}

// the length of cards must be 2
func _RecordWinRateForSpecificHand(cards []Card, winRate float64, winCount int, trials int) error {
	if len(cards) != 2 {
		return errors.New("the length of cards must be 2")
	}
	filename := "win_rate.txt"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("error opening file: ", err)
	}
	defer file.Close()

	cards_str, _ := ConvertCardsToSymbol(cards)
	now := time.Now().Format("2006-01-02 15:04:05")
	content := fmt.Sprintf("%s %s %.4f %d %d", now, cards_str, winRate, winCount, trials)
	if _, err := file.WriteString(content + "\n"); err != nil {
		return err
	} else {
		return nil
	}
}

func CalculateWinRateForSpecificHand(cards []Card) {
	deck := NewDeck()
	cards_p1, _ := deck.DrawSpecificCards(cards)
	deck.Shuffle()

	trials := 100000
	winCount := 0
	chopCount := 0
	loseCount := 0
	for i := 0; i < trials; i++ {
		cards_p2, _ := deck.DrawCards(2)
		board, _ := deck.DrawCards(5)
		cards_p1_board := append(cards_p1, board...)
		cards_p2_board := append(cards_p2, board...)
		hand_p1, _ := DetermineHandRank(cards_p1_board)
		hand_p2, _ := DetermineHandRank(cards_p2_board)

		resultCompareTwoHands, _ := CompareTwoHands(hand_p1, hand_p2)
		if resultCompareTwoHands == 1 {
			winCount++
		} else if resultCompareTwoHands == 0 {
			chopCount++
		} else if resultCompareTwoHands == 2 {
			loseCount++
		}
		deck.AddCards(cards_p2_board)
		deck.Shuffle()

		// Print the result of each trial
		// fmt.Println("Trial", i+1, ":", "Player 1's cards:", cards_p1, "Player 2's cards:", cards_p2, "Board cards:", board, "Player", resultCompareTwoHands, "wins")
	}
	winRate := float64(winCount) / float64(trials) * 100
	fmt.Printf("Win rate: %.4f%% Wins: %d Loses: %d Chops: %d Trials: %d\n", winRate, winCount, loseCount, chopCount, trials)
	_RecordWinRateForSpecificHand(cards, winRate, winCount, trials)
}

func main() {
	cards := []Card{{Suit: "s", Number: 1}, {Suit: "s", Number: 13}}
	CalculateWinRateForSpecificHand(cards)
}

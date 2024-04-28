package main

import (
	"errors"
	"fmt"
	"math/rand"
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

type Position int

const (
	SB Position = iota
	BB
)

type Hand struct {
	pid   int
	Rank  HandRank
	Cards []Card
}

type Player struct {
	Id         int
	Position   Position
	Stuck      int
	Cards      []Card
	Betting    int
	LastAction *string
}

type Board struct {
	Cards   []Card
	Pod     int
	BetSize int
}

func DetermineHandRank(cards []Card) (Hand, error) {
	if len(cards) < 5 {
		return Hand{}, errors.New("given cards are less than 5 and no rank can be determined")
	}

	// Initialize hand
	// fmt.Println("Cards:", cards)
	rank := JudgeHandRank(cards)
	// fmt.Println("Rank:", rank, "Cards:", cards)

	selectedCards := SelectStrongestCards(cards, rank)

	hand := Hand{Rank: rank, Cards: selectedCards}
	// fmt.Println("Cards: ", cards, "Hand:", hand.Rank, hand.Cards)
	return hand, nil
}

// Warnings: Length of hand1.Cards and hand2.Cards must be 5
// Return 1 if hand1 wins, 2 if hand2 wins, 0 if it's a tie, -1 if error
func CompareTwoHands(hand1 Hand, hand2 Hand) (int, error) {
	if len(hand1.Cards) != 5 {
		fmt.Println("hand1.Cards:", hand1.Cards)
		return -1, errors.New("length of hand1.Cards must be 5")
	} else if len(hand2.Cards) != 5 {
		fmt.Println("hand2.Cards:", hand2.Cards)
		return -1, errors.New("length of hand2.Cards must be 5")
	}
	// Return 1 if hand1 wins, 2 if hand2 wins, 0 if it's a tie
	if hand1.Rank > hand2.Rank {
		return 1, nil
	} else if hand1.Rank < hand2.Rank {
		return 2, nil
	}

	// Compare the cards of the same rank
	cards_p1 := ReplaceAceWithFouteenFromOne(hand1.Cards)
	cards_p2 := ReplaceAceWithFouteenFromOne(hand2.Cards)
	for i := 0; i < 5; i++ {
		if cards_p1[i].Number > cards_p2[i].Number {
			return 1, nil
		} else if cards_p1[i].Number < cards_p2[i].Number {
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
	case 10:
		return "T"
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

func main() {
	deck := NewDeck()
	deck.Shuffle()

	action_list := []string{
		"f",
		"cK", "cRf", "cRc", "cRrF", "cRrC", "cRrRf", "cRrRc", "cRrRaF", "cRrRaC",
		"rF", "rC", "rRf", "rRc", "rRrF", "rRrC", "rRrAf", "rRrAc",
	}

	f := 0
	c := 0
	r := 0
	trials := 1000000
	result := make(map[string]float64)
	for i := 0; i < trials; i++ {
		// Create Player 1
		player_1 := Player{Id: 1, Stuck: 100, Position: SB, LastAction: nil}
		cards_p1, _ := deck.DrawSpecificCards([]Card{{Number: 1, Suit: "s"}, {Number: 13, Suit: "d"}})
		player_1.Cards = cards_p1
		// fmt.Println("Player 1's cards:", cards_p1)

		// Create Player 2
		player_2 := Player{Id: 2, Stuck: 100, Position: BB, LastAction: nil}
		cards_p2, _ := deck.DrawCards(2)
		player_2.Cards = cards_p2

		// Init Board
		board := Board{Pod: 0, BetSize: 0}
		bb := 2
		sb := 1
		player_1.Betting = sb
		player_1.Stuck -= sb
		player_2.Betting = bb
		player_2.Stuck -= bb
		board.BetSize = bb

		// Determine Actions
		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(len(action_list))
		actions := action_list[randomIndex]

		for _, action_chr := range actions {
			if action_chr == 'f' {
				// Put the chips which put out into the pod
				board.Pod += player_1.Betting
				player_1.Betting = 0
				// Get the pod and return the chips which put out
				player_2.Stuck += board.Pod + player_2.Betting
				board.Pod = 0
				player_2.Betting = 0
				// Player 1 Action Done
				player_1.LastAction = &actions
				break
			} else if action_chr == 'c' {
				// Subtract the difference between the raise size and the current betting size from the player's stuck
				player_1.Stuck -= board.BetSize - player_1.Betting
				player_1.Betting = board.BetSize
				// Player 2 Action Done
				player_1.LastAction = &actions
			} else if action_chr == 'r' {
				// Determine raise size
				raise_size := board.BetSize * 3
				// Subtract the difference between the raise size and the current betting size from the player's stuck
				player_1.Stuck -= raise_size - player_1.Betting
				player_1.Betting = raise_size
				board.BetSize = raise_size
				// Player 1 Action Done
				player_1.LastAction = &actions
			} else if action_chr == 'a' {
				// Determine raise size
				raise_size := player_1.Stuck + player_1.Betting
				// Subtract the difference between the raise size and the current betting size from the player's stuck
				player_1.Stuck = 0
				player_1.Betting = raise_size
				board.BetSize = raise_size
			} else if action_chr == 'F' {
				// Put the chips which put out into the pod
				board.Pod += player_2.Betting
				player_2.Betting = 0
				// Get the pod and return the chips which put out
				player_1.Stuck += board.Pod + player_1.Betting
				board.Pod = 0
				player_1.Betting = 0
				// Player 2 Action Done
				player_2.LastAction = &actions
				break
			} else if action_chr == 'C' {
				// Subtract the difference between the raise size and the current betting size from the player's stuck
				player_2.Stuck -= board.BetSize - player_2.Betting
				player_2.Betting = board.BetSize
				// Player 2 Action Done
				player_2.LastAction = &actions
			} else if action_chr == 'R' {
				// Determine raise size
				raise_size := board.BetSize * 3
				// Subtract the difference between the raise size and the current betting size from the player's stuck
				player_2.Stuck -= raise_size - player_2.Betting
				player_2.Betting = raise_size
				board.BetSize = raise_size
				// Player 2 Action Done
				player_2.LastAction = &actions
			} else if action_chr == 'K' {
				board.Pod += player_1.Betting + player_2.Betting
				player_1.Betting = 0
				player_2.Betting = 0
				// Player 2 Action Done
				player_2.LastAction = &actions
			} else if action_chr == 'A' {
				// Determine raise size
				raise_size := player_2.Stuck + player_2.Betting
				// Subtract the difference between the raise size and the current betting size from the player's stuck
				player_2.Stuck = 0
				player_2.Betting = raise_size
				board.BetSize = raise_size
			}

			// Determine Showdown
			showdown := false
			if player_1.LastAction != nil && player_2.LastAction != nil && player_1.Betting == player_2.Betting {
				showdown = true
			}

			if showdown {
				// Put all bets in the pot
				board.Pod += player_1.Betting + player_2.Betting
				// Decide the winner
				board.Cards, _ = deck.DrawCards(5)
				hand_p1, _ := DetermineHandRank(append(player_1.Cards, board.Cards...))
				hand_p2, _ := DetermineHandRank(append(player_2.Cards, board.Cards...))
				resultCompareTwoHands, _ := CompareTwoHands(hand_p1, hand_p2)
				if resultCompareTwoHands == 1 {
					player_1.Stuck += board.Pod
				} else if resultCompareTwoHands == 2 {
					player_2.Stuck += board.Pod
				} else if resultCompareTwoHands == 0 {
					player_1.Stuck += board.Pod / 2
					player_2.Stuck += board.Pod / 2
				} else {
					panic("error -1")
				}
			}
		}
		deck.AddCards(board.Cards)
		deck.AddCards(player_1.Cards)
		deck.AddCards(player_2.Cards)
		deck.Shuffle()
		res := float64(player_1.Stuck) - 100
		result[actions] += res
		if actions[0] == 'f' {
			f++
		} else if actions[0] == 'c' {
			c++
		} else if actions[0] == 'r' {
			r++
		}
	}

	first_action_result := make(map[string]float64)
	for key, value := range result {
		first_action := string(key[0])
		first_action_result[first_action] += value
	}
	for key, value := range first_action_result {
		ev := 0.00
		if key == "f" {
			ev = value / float64(f)
		} else if key == "c" {
			ev = value / float64(c)
		} else if key == "r" {
			ev = value / float64(r)
		}
		fmt.Printf("First Action: %s, Result: %.2f\n", key, ev)
	}
}

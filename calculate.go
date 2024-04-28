package main

import (
	"fmt"
	"time"
)

func CalculateWinRateForSpecificHand(cards []Card, trials int) {
	deck := NewDeck()
	cards_p1, _ := deck.DrawSpecificCards(cards)
	deck.Shuffle()

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

		resultCompareTwoHands, err := CompareTwoHands(hand_p1, hand_p2)
		if resultCompareTwoHands == 1 {
			winCount++
		} else if resultCompareTwoHands == 0 {
			chopCount++
		} else if resultCompareTwoHands == 2 {
			loseCount++
		} else {
			fmt.Println("Error:", err)
		}
		deck.AddCards(cards_p2_board)
		deck.Shuffle()
	}
	winRate := float64(winCount) / float64(trials) * 100
	// fmt.Printf("Win rate: %.4f%% Wins: %d Loses: %d Chops: %d Trials: %d\n", winRate, winCount, loseCount, chopCount, trials)
	err := RecordWinRateForSpecificHand(cards, winRate, winCount, loseCount, chopCount, trials)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func CalculateWinRate() {
	deck_org := NewDeck()
	for _, card_i := range deck_org.Cards {
		for _, card_j := range deck_org.Cards {
			if card_i == card_j {
				continue
			}
			cards := []Card{card_i, card_j}
			CalculateWinRateForSpecificHand(cards, 1000)
		}
	}
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

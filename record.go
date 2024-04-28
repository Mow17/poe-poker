package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// Warnings: the length of cards must be 2
func RecordWinRateForSpecificHand(cards []Card, winRate float64, winCount int, loseCount int, chopCount int, trials int) error {
	if len(cards) != 2 {
		return errors.New("the length of cards must be 2")
	}
	filename := "output/log/log_win_rate.txt"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("error opening file: ", err)
	}
	defer file.Close()

	cards_str, _ := ConvertCardsToSymbol(cards)
	now := time.Now().Format("2006-01-02 15:04:05")
	content := fmt.Sprintf("%s %s %.4f %d %d %d %d", now, cards_str, winRate, winCount, loseCount, chopCount, trials)
	if _, err := file.WriteString(content + "\n"); err != nil {
		return err
	} else {
		return nil
	}
}

func RecordAggregateWinRate(key string, winRate float64, now string) error {
	filename := "output/analysis/aggregated_win_rate_" + now + ".txt"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("error opening file: ", err)
	}
	defer file.Close()

	content := fmt.Sprintf("%s %.4f", key, winRate)
	if _, err := file.WriteString(content + "\n"); err != nil {
		return err
	} else {
		return nil
	}
}

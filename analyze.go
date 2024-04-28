package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

type WinCounter struct {
	winCount int
	trials   int
}

func AggregateRandomHandWinRate() error {
	inputFilename := "output/log/log_win_rate.txt"

	// Read the file
	inputFile, err := os.Open(inputFilename)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	winCounter := make(map[string]WinCounter)

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		key := fields[2]
		winCount, _ := strconv.Atoi(fields[4])
		trials, _ := strconv.Atoi(fields[7])
		winCounter[key] = WinCounter{winCounter[key].winCount + winCount, winCounter[key].trials + trials}
	}

	now := time.Now().Format("2006-01-02T15:04:05")
	for key, value := range winCounter {
		err = RecordAggregateWinRate(key, float64(value.winCount)/float64(value.trials), now)
		if err != nil {
			return err
		}
	}
	return nil
}

package main

func JudgeHandRank(cards_org []Card) HandRank {
	cards := cards_org
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
	}
	return handRank
}

func JudgeStraightFlush(cards []Card) bool {
	cards = AddFourteen(cards)
	cards = SortDecreaseingSuitPriority(cards)
	straightFlush := false

	tmp_count := 1
	tmp_card := cards[0]
	for i, card := range cards {
		if i == 0 {
			continue
		}
		if card.Number == tmp_card.Number-1 && card.Suit == tmp_card.Suit {
			tmp_count++
			if tmp_count == 5 {
				straightFlush = true
				break
			}
		} else {
			tmp_count = 1
		}
		tmp_card = card
	}
	RemoveFourteen(cards)
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
		if !threeOfAKind && count == 3 {
			threeOfAKind = true
		} else if count >= 2 {
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

	previousCard := cards[0]
	count := 1
	for i, card := range cards {
		if i == 0 {
			continue
		}
		if card.Number == previousCard.Number-1 {
			previousCard = card
			count++
			if count == 5 {
				straight = true
				break
			}
		} else if card.Number == previousCard.Number {
			continue
		} else {
			previousCard = card
			count = 1
		}
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

	if pairCount >= 2 {
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

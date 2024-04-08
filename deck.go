package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	Suit   string
	Number int
}

type Deck struct {
	Cards []Card
}

// Create a deck of 52 playing cards, excluding jokers
// s: Spade, h: Heart, d: Diamond, c: Clover
func NewDeck() Deck {
	suits := []string{"s", "h", "d", "c"}
	cards := []Card{}

	for _, suit := range suits {
		for i := 1; i <= 13; i++ {
			card := Card{Suit: suit, Number: i}
			cards = append(cards, card)
		}
	}

	return Deck{Cards: cards}
}

// Standard output in 4rows over 13 columns
func (d Deck) Display() {
	for i, card := range d.Cards {
		if i%13 != 0 {
			fmt.Printf(" ")
		}
		if card.Number < 10 {
			fmt.Printf(" ")
		}
		fmt.Printf("%s%d", card.Suit, card.Number)
		if (i+1)%13 == 0 || i+1 == len(d.Cards) {
			fmt.Printf("\n")
		}
	}
}

// Shuffle using Fisher-Yates algorithm
func (d *Deck) Shuffle() {
	deck_size := len(d.Cards)

	// Change seed based on curret time
	rand.Seed(time.Now().UnixNano())

	for i := deck_size - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}

func (d *Deck) DrawCards(numCards int) ([]Card, error) {
	if len(d.Cards) < numCards {
		return nil, errors.New("error: Not enough cards in the deck")
	}
	drawnCards := d.Cards[:numCards]
	// Remove drawn cards from deck
	d.Cards = d.Cards[numCards:]

	return drawnCards, nil
}

func (d *Deck) DrawSpecificCards(cards []Card) ([]Card, error) {
	for _, card := range cards {
		found := false
		for i, deckCard := range d.Cards {
			if card.Suit == deckCard.Suit && card.Number == deckCard.Number {
				d.Cards = append(d.Cards[:i], d.Cards[i+1:]...)
				found = true
				break
			}
		}
		if !found {
			return nil, errors.New("error: Card not found in the deck")
		}
	}
	return cards, nil
}

func (d *Deck) AddCards(cards []Card) {
	d.Cards = append(d.Cards, cards...)
}

func (d *Deck) ValidateDeck() {
	uniqueCards := make(map[Card]bool)
	for _, card := range d.Cards {
		if _, ok := uniqueCards[card]; ok {
			panic("error: Duplicate cards in the deck")
		}
		uniqueCards[card] = true
	}
}

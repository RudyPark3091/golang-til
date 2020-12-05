package main

import (
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
}

func NewDeck() *Deck {
	deck := &Deck{}

	for i := 0; i < 52; i++ {
		switch {
		case i < 13:
			deck.Cards = append(deck.Cards, Card{i%13 + 1, SPADE})
		case i < 26:
			deck.Cards = append(deck.Cards, Card{i%13 + 1, CLOVER})
		case i < 39:
			deck.Cards = append(deck.Cards, Card{i%13 + 1, HEART})
		case i < 52:
			deck.Cards = append(deck.Cards, Card{i%13 + 1, DIAMOND})
		}
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(52, func(i, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})

	return deck
}

func (d *Deck) Draw() Card {
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card
}

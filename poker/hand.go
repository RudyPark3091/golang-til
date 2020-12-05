package main

import (
	"sort"
)

type rank int

const (
	HIGH_CARD rank = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	STRAIGHT
	ROYAL_STRAIGHT
	FLUSH
	FULL_HOUSE
	FOUR_OF_A_KIND
	STRAIGHT_FLUSH
	ROYAL_STRAIGHT_FLUSH
)

var RankToString = map[rank]string{
	0:  "High Card",
	1:  "One Pair",
	2:  "Two Pair",
	3:  "Three of a Kind",
	4:  "Straight",
	5:  "Royal Straight",
	6:  "Flush",
	7:  "Full House",
	8:  "Four of a Kind",
	9:  "Straight Flush",
	10: "Royal Straight Flush",
}

type Hand struct {
	Cards []Card
}

func (h *Hand) Draw(d *Deck) {
	h.Cards = append(h.Cards, d.Draw())
}

func (by By) Sort(c []Card) {
	cs := &CardSorter{
		Cards: c,
		by:    by,
	}
	sort.Sort(cs)
}

type By func(c1, c2 *Card) bool

type CardSorter struct {
	Cards []Card
	by    By
}

func (s *CardSorter) Len() int {
	return len(s.Cards)
}

func (s *CardSorter) Swap(i, j int) {
	s.Cards[i], s.Cards[j] = s.Cards[j], s.Cards[i]
}

func (s *CardSorter) Less(i, j int) bool {
	return s.by(&s.Cards[i], &s.Cards[j])
}

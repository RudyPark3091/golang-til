package main

import (
	"fmt"
)

type HandInfo struct {
	Rank     int
	RankName string
	Top      int
	FullName string
}

func ShowCardsInHand(h *Hand) {
	fmt.Println("hand:")
	for _, c := range h.Cards {
		fmt.Print(c)
	}
	fmt.Println()
}

func InitializeInfoStruct(rank rank, top Card) *HandInfo {
	info := new(HandInfo)
	info.Rank = int(rank)
	info.RankName = RankToString[rank]
	info.Top = top.Number
	info.FullName = ShapeToString[top.Shape] +
		" " + NumberToName[top.Number] +
		" " + RankToString[rank]

	return info
}

func EvaluatePairs(h *Hand) (bool, *Card, int) {
	hasPair := false
	pairTop := &Card{}
	pairCount := 0

	for i := 1; i < len(h.Cards); i++ {
		if h.Cards[i].Number == h.Cards[i-1].Number {
			hasPair = true
			pairTop = &h.Cards[i]
			pairCount++
		}
	}

	return hasPair, pairTop, pairCount
}

func EvaluateTriple(h *Hand) (bool, *Card) {
	hasTriple := false
	tripleTop := &Card{}

	for i := 2; i < len(h.Cards); i++ {
		yes := h.Cards[i].Number == h.Cards[i-1].Number &&
			h.Cards[i-1].Number == h.Cards[i-2].Number

		if yes {
			hasTriple = true
			tripleTop = &h.Cards[i]
			h.Cards = append(h.Cards[:i-2], h.Cards[i+1:]...)
		}
	}

	return hasTriple, tripleTop
}

func EvaluateFourCards(h *Hand) (bool, *Card) {
	hasFour := h.Cards[0].Number == h.Cards[3].Number
	hasFour2 := h.Cards[1].Number == h.Cards[4].Number

	if hasFour {
		return hasFour, &h.Cards[3]
	}
	if hasFour2 {
		return hasFour2, &h.Cards[4]
	}

	return false, &Card{}
}

func EvaluateStraightFlush(h *Hand) (bool, bool) {
	hasStraight := false
	hasFlush := false

	hasStraight = h.Cards[1].Number == h.Cards[0].Number+1 &&
		h.Cards[2].Number == h.Cards[1].Number+1 &&
		h.Cards[3].Number == h.Cards[2].Number+1 &&
		h.Cards[4].Number == h.Cards[3].Number+1

	hasFlush = h.Cards[0].Shape == h.Cards[1].Shape &&
		h.Cards[1].Shape == h.Cards[2].Shape &&
		h.Cards[2].Shape == h.Cards[3].Shape &&
		h.Cards[3].Shape == h.Cards[4].Shape

	return hasStraight, hasFlush
}

func EvaluateRoyalStraight(h *Hand) bool {
	isRoyalStraight := false

	isRoyalStraight = h.Cards[0].Number == ACE &&
		h.Cards[1].Number == TEN &&
		h.Cards[2].Number == JACK &&
		h.Cards[3].Number == QUEEN &&
		h.Cards[4].Number == KING

	return isRoyalStraight
}

func Evaluate(h *Hand) *HandInfo {
	info := new(HandInfo)

	num := func(c1, c2 *Card) bool {
		return c1.Number < c2.Number
	}
	By(num).Sort(h.Cards)

	isRoyalStraight := EvaluateRoyalStraight(h)

	hasStraight, hasFlush := EvaluateStraightFlush(h)
	if isRoyalStraight {
		if hasFlush {
			info = InitializeInfoStruct(ROYAL_STRAIGHT_FLUSH, h.Cards[4])
			return info
		}

		if !hasFlush {
			info = InitializeInfoStruct(ROYAL_STRAIGHT, h.Cards[4])
			return info
		}
	} else {
		if hasFlush && hasStraight {
			info = InitializeInfoStruct(STRAIGHT_FLUSH, h.Cards[4])
			return info
		}

		if hasFlush && !hasStraight {
			info = InitializeInfoStruct(FLUSH, h.Cards[4])
			return info
		}

		if hasStraight && !hasFlush {
			info = InitializeInfoStruct(STRAIGHT, h.Cards[4])
			return info
		}
	}

	hasFour, fourTop := EvaluateFourCards(h)
	if hasFour {
		info = InitializeInfoStruct(FOUR_OF_A_KIND, *fourTop)
		return info
	}

	hasTriple, tripleTop := EvaluateTriple(h)
	if hasTriple {
		info = InitializeInfoStruct(THREE_OF_A_KIND, *tripleTop)
	}

	hasPair, pairTop, pairCount := EvaluatePairs(h)
	if hasTriple && hasPair {
		// Full House
		info = InitializeInfoStruct(FULL_HOUSE, *tripleTop)
		return info
	}

	if hasPair {
		if pairCount == 1 {
			info = InitializeInfoStruct(ONE_PAIR, *pairTop)
		} else {
			info = InitializeInfoStruct(TWO_PAIR, *pairTop)
		}
	}

	if !hasPair {
		info = InitializeInfoStruct(HIGH_CARD, h.Cards[4])
	}

	return info
}

func main() {
	deck := NewDeck()

	h1 := &Hand{}
	h2 := &Hand{}
	h3 := &Hand{[]Card{
		{1, 0},
		{10, 0},
		{11, 0},
		{12, 0},
		{13, 0},
	}}
	h4 := &Hand{[]Card{
		{1, 0},
		{1, 1},
		{1, 2},
		{3, 3},
		{3, 1},
	}}

	for i := 0; i < 5; i++ {
		h1.Draw(deck)
		h2.Draw(deck)
	}

	info1 := Evaluate(h1)
	info2 := Evaluate(h2)
	info3 := Evaluate(h3)
	info4 := Evaluate(h4)

	ShowCardsInHand(h1)
	ShowCardsInHand(h2)
	ShowCardsInHand(h3)
	ShowCardsInHand(h4)

	fmt.Printf("%+v\n", info1)
	fmt.Printf("%+v\n", info2)
	fmt.Printf("%+v\n", info3)
	fmt.Printf("%+v\n", info4)
}

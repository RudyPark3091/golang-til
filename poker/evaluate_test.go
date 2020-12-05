package main

import (
	"testing"
)

var num = func(c1, c2 *Card) bool {
	return c1.Number < c2.Number
}

func TestEvaluatePairs(t *testing.T) {
	// Ace One Pair
	h := &Hand{[]Card{
		{ACE, SPADE},
		{EIGHT, CLOVER},
		{ACE, HEART},
		{QUEEN, DIAMOND},
		{KING, CLOVER},
	}}

	By(num).Sort(h.Cards)
	hasPair, pairTop, pairCount := EvaluatePairs(h)

	if !hasPair {
		t.Error("Cannot Detect One Pair")
	}

	ok := *pairTop == Card{ACE, SPADE} || *pairTop == Card{ACE, HEART}
	if !ok {
		t.Error("Cannot Detect Pair Top")
	}

	if pairCount != 1 {
		t.Error("Cannot Detect Pair Count")
	}
}

func TestEvaluateTriple(t *testing.T) {
	// Three Three of Kind
	h := &Hand{[]Card{
		{THREE, SPADE},
		{ACE, HEART},
		{JACK, HEART},
		{THREE, DIAMOND},
		{THREE, CLOVER},
	}}

	By(num).Sort(h.Cards)
	hasTriple, tripleTop := EvaluateTriple(h)

	if !hasTriple {
		t.Error("Cannot Detect Triple")
	}

	ok := *tripleTop == Card{THREE, SPADE} ||
		*tripleTop == Card{THREE, DIAMOND} ||
		*tripleTop == Card{THREE, CLOVER}
	if !ok {
		t.Error("Cannot Detect Three of a Kind Top")
	}
}

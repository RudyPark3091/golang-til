package main

type shape int

const (
	SPADE shape = iota
	CLOVER
	HEART
	DIAMOND
)

const (
	_ int = iota
	ACE
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
)

var NumberToName = map[int]string{
	1:  "Ace",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
	11: "Jack",
	12: "Queen",
	13: "King",
}

var ShapeToString = map[shape]string{
	0: "Spade",
	1: "Clover",
	2: "Heart",
	3: "Diamond",
}

type Card struct {
	Number int
	Shape  shape
}

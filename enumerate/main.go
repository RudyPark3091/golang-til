package main

import (
	"fmt"
)

type snack int

const HOTDOG snack = 0
const CHOCOLATE snack = 1
const ICECREAM snack = 2

const (
	FRENCHFRIES snack = 3
	BISQUIT     snack = 4
	CHEESE      snack = 5
)

const (
	CAKE snack = iota
	NOODLE
	CEREAL
)

type ByteSize int

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

func main() {
	fmt.Printf("%d\n", HOTDOG)      // 0
	fmt.Printf("%d\n", CHOCOLATE)   // 1
	fmt.Printf("%d\n", ICECREAM)    // 2
	fmt.Printf("%d\n", FRENCHFRIES) // 3
	fmt.Printf("%d\n", BISQUIT)     // 4
	fmt.Printf("%d\n", CHEESE)      // 5
	fmt.Printf("%d\n", CAKE)        // 0
	fmt.Printf("%d\n", NOODLE)      // 1
	fmt.Printf("%d\n", CEREAL)      // 2

	fmt.Printf("ByteSize %d Byte\n", KB)
	fmt.Printf("ByteSize %d Byte\n", MB)
	fmt.Printf("ByteSize %d Byte\n", GB)
	fmt.Printf("ByteSize %d Byte\n", TB)
}

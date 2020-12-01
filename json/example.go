package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Color struct {
	Color string `json:"color"`
	Value string `json:"value"`
}

func Example() {
	f, err := os.Open("./example.json")
	if err != nil {
		panic(err)
	}

	colors := new([]Color)
	d := json.NewDecoder(f)
	d.Decode(colors)

	fmt.Printf("%+v\n", colors)
}

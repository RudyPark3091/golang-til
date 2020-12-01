package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Item struct {
	Id       string    `json:"id"`
	Type     string    `json:"type"`
	Name     string    `json:"name"`
	Ppu      int       `json:"ppu"`
	Batters  Batter    `json:"batters"`
	Toppings []Topping `json:"topping"`
}

type Batter struct {
	BatterItems []BatterItem `json:"batter"`
}

type BatterItem struct {
	Id   string `json:"id"`
	Type string `json:"Type"`
}

type Topping struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

func ExampleTwo() {
	f, err := os.Open("./example2.json")
	if err != nil {
		panic(err)
	}

	item := new(Item)
	d := json.NewDecoder(f)
	d.Decode(item)

	fmt.Printf("%+v\n", item)
}

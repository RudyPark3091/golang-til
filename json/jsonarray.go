package main

import (
	"encoding/json"
	"fmt"
)

type Recipe struct {
	Name       string   `json:"name"`
	Ingredient []string `json:"ingredient"`
}

func JsonArray() {
	cake := &Recipe{
		"cake",
		[]string{"sugar", "flour", "eggs"},
	}

	cakeJson, err := json.Marshal(cake)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(cakeJson))
}

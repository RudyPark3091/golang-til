package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("text.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("File Content")
	fmt.Println(strings.Repeat("=", 30))
	// must be formatted as string
	fmt.Printf("%s", content)
}

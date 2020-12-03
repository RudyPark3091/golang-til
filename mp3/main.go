package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// ID3v1 TAG processing
func main() {
	// reading file
	data, err := ioutil.ReadFile("sample.mp3")
	if err != nil {
		panic(err)
	}

	// opening file stream
	f, err := os.Create("sample2.mp3")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// making TAG byte slice
	tag := []byte("TAG")

	title := []byte("Hello0000000000000000000000000")
	artist := []byte("artist000000000000000000000000")
	album := []byte("album0000000000000000000000000")
	year := []byte("2020")
	comment := []byte("comment00000000000000000000000")
	genre := byte(0)

	tag = append(tag, title...)
	tag = append(tag, artist...)
	tag = append(tag, album...)
	tag = append(tag, year...)
	tag = append(tag, comment...)
	tag = append(tag, genre)

	fmt.Println(len(string(tag)))

	// appending ID3 TAG
	data = append(data, tag...)

	r := bytes.NewReader(data)

	// write to a file stream
	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}
}

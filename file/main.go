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

	fmt.Println(strings.Repeat("=", 15), "Downloading Image", strings.Repeat("=", 15))

	url := "https://images.unsplash.com/photo-1602526216007-4a4a197b44aa?ixlib=rb-1.2.1&ixid=MXwxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHwxfHx8ZW58MHx8fA%3D%3D&auto=format&fit=crop&w=500&q=60"
	filename := "sample.png"
	targetPath := "./image"
	// written in ./image-dl.go

	err = DownloadImage(url, filename, targetPath)
	if err != nil {
		panic(err)
	}
}

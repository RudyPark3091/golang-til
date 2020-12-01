package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Ide struct {
	Lang    string `json:"JsonLang"`
	IdeName string `json:"JsonIdeName"`
}

// using Marshal()
func StructToJson() {
	py := &Ide{"Python", "PyCharm"}

	// json.Marshal()
	// Golang struct to JSON format
	// returns byte data -> string(data) will show readable string form
	data, err := json.Marshal(py)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	// json.MarshalIndent(v interface{}, prefix, indent string)
	// same with Marshal() but with indentation
	data, err = json.MarshalIndent(py, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

// using Unmarshal()
func JsonToStruct() {
	// reading JSON file
	data, err := ioutil.ReadFile("./data.json")
	if err != nil {
		panic(err)
	}

	// creating new struct variable
	ide := new(Ide)
	// mapping to struct
	err = json.Unmarshal(data, ide)
	if err != nil {
		panic(err)
	}

	// data access
	fmt.Println("lang:", ide.Lang)
	fmt.Println("ide:", ide.IdeName)

	fmt.Printf("%+v\n", ide)
}

func EncodingJson() {
	e := json.NewEncoder(os.Stdout)
	java := &Ide{"java", "IntelliJ"}

	e.Encode(java)
}

func EncodeAndWriteFile() {
	f, err := os.Create("encoded.json")
	if err != nil {
		panic(err)
	}

	golang := &Ide{"golang", "goland"}

	e := json.NewEncoder(f)
	e.SetIndent("", "	")
	e.Encode(golang)
}

func DecodingJson() {
	ide := new(Ide)
	f, err := os.Open("./data.json")
	if err != nil {
		panic(err)
	}

	d := json.NewDecoder(f)
	d.Decode(&ide)

	fmt.Printf("%+v\n", ide)
}

func main() {
	StructToJson()
	JsonToStruct()
	EncodingJson()
	EncodeAndWriteFile()
	DecodingJson()
	JsonArray()
	Example()
	ExampleTwo()
}

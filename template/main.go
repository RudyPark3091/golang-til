package main

import (
	"fmt"
	"os"
	"text/template"
)

type Num struct {
	Number int
}

func main() {
	number := Num{3}
	tmpl, err := template.New("test").Parse("the number is {{.Number}}")
	tmpl.Execute(os.Stdout, number)
	tmpl, err = template.New("cat").Parse("this is cat {{.Number}}")
	if err != nil {
		panic(err)
	}

	tmpl.Lookup("cat")
	tmpl.Execute(os.Stdout, number)
	fmt.Println()
}

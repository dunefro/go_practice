package main

import (
	"os"
	"text/template"
)

const templ = `{{ range . }}Hello {{ . }}
{{ end }}`

type output string

func main() {
	names := []string{"vedant", "harry", "ron", "hagrid", "tom"}
	t, err := template.New("hogwarts").Parse(templ)
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdin, names)
	if err != nil {
		panic(err)
	}
}

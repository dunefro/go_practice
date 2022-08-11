package main

import (
	"os"
	"text/template"
)

const templ = `{{ range .}}{{ .}}
{{end}}`

func main() {
	t, _ := template.New("person").Parse(templ)

	t.Execute(os.Stdin, []string{"a", "b", "c", "d"})
}

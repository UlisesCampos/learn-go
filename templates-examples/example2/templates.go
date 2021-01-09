package templates

import (
	"html/template"
	"os"
	"strings"
)

const sampleTemplate = ` 
	This template demonstrates a {{ .Variable | printf "%#v" }}.
	{{if .Condition}}
	If condition is set, we'll print this
	{{else}} Oterwisem weĺl print this instead
	{{end}}
	Next we'll iterate over an array of strings:
	{{range $index, $item := .Items}}
	{{$index}}: {{$item}}
	{{end}}

	We can also easily import other functions like strings.Split
	then inmediately used the array created as a result:
	{{ range $index, $item := split .Words ","}}
	{{$index}}: {{$item}}
	{{end}}
	Blocks are a way to embed templates into one another
	{{ block "block_example" .}}
	No block defined!
	{{end}}
	{{/*
	This is a way to insert
	a multi-line comment 	
	*/}}
`

const secondTemplate = `
	{{ define "block_example"}}
	{{.OtherVariable}}
	{{end}}
`

// RunTemplate initializes a template and demostrates a variety of template helper functions
func RunTemplate() error {
	data := struct {
		Condition     bool
		Variable      string
		Items         []string
		Words         string
		OtherVariable string
	}{
		Condition:     true,
		Variable:      "variable",
		Items:         []string{"item1", "item2", "item3"},
		Words:         "words",
		OtherVariable: "I defined in a second template!",
	}
	funcmap := template.FuncMap{
		"split": strings.Split,
	}
	// these can also be chained
	t := template.New("example")
	t = t.Funcs(funcmap)
	// We could use Must instead to panic on error template.Must(t.Parse(sampleTemplate))
	t, err := t.Parse(sampleTemplate)
	if err != nil {
		return err
	}
	// To demostrates blocks we'll create another template by cloning the
	// first template, then parsing a second
	t2, err := t.Clone()
	if err != nil {
		return err
	}
	t2, err = t2.Parse(secondTemplate)
	if err != nil {
		return err
	}
	// write the template to stdout and populate it with data
	err = t2.Execute(os.Stdout, &data)
	if err != nil {
		return err
	}
	return nil
}

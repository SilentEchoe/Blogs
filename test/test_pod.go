package main

import (
	"os"
	"text/template"
)

func main() {
	// Define a map with some key-value pairs
	myMap := map[string]interface{}{
		"name":      "John",
		"age":       30,
		"isMarried": true,
	}

	// Use Go template to access the map values with if condition
	t := template.Must(template.New("myTemplate").Parse(`
    {{if eq .name "John"}}
        {{if .isMarried}}
            {{.name}} is married
        {{else}}
            {{.name}} is not married
        {{end}}
    {{else}}
        {{.name}} is not John
    {{end}}
`))

	// Execute the template with the map as input
	err := t.Execute(os.Stdout, myMap)
	if err != nil {
		panic(err)
	}

}

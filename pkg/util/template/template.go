package template

import (
	"bytes"
	"html/template"
)

// Templater : a templating mini engine that can tempate html files
type Templater struct {
	template *template.Template
}

// New : Generate a new templater with a file path
func New(filePathToTemplate string) *Templater {
	return &Templater{
		template: template.Must(template.ParseFiles(filePathToTemplate)),
	}
}

// Generate string representation of the template with error state
func (t *Templater) Generate(data interface{}) (string, error) {
	buf := &bytes.Buffer{}
	err := t.template.Execute(buf, data)
	return buf.String(), err
}

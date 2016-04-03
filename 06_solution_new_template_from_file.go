package main

import (
	"fmt"
	"io/ioutil"

	"text/template"
)

func NewTemplateFromFile(fileName string) (*template.Template, error) {
	src, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("error parsing static/index.html: %s", err)
	}

	t, err := template.New("index.html").Parse(string(src))
	if err != nil {
		return nil, fmt.Errorf("error creating template for index.html: %s", err)
	}

	return t, nil
}

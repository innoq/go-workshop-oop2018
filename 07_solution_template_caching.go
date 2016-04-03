package main

import (
	"log"
	"text/template"
)

// copy to main.go
var indexTemplate *template.Template

func init() {
	var err error
	indexTemplate, err = NewTemplateFromFile("static/index.html")
	if err != nil {
		log.Fatal("error creating template for index.html: %s", err)
	}
}

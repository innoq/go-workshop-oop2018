package main

import (
	"os"
	"testing"
)

func TestNewTemplateFromFile(t *testing.T) {
	_, err := NewTemplateFromFile("nonexistant")
	if err == nil {
		t.Errorf("NewTemplateFromFile schlägt nicht fehl, wenn es eine nicht existente Datei bekommt")
	}

	// {{ .
	_, err = NewTemplateFromFile("./06_test_new_template_from_file_test.go")
	if err == nil {
		t.Errorf("NewTemplateFromFile schlägt nicht fehl, wenn ein defektes Template eingelesen wird")
	}

}

func ExampleNewTemplateFromFile() {
	t, _ := NewTemplateFromFile("./06_test_template.txt")
	t.Execute(os.Stdout, struct{ Foo string }{"bar"})

	// Output: bar
}
